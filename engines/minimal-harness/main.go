// Command minimal-harness is the smallest agent harness that is still useful for
// real coding work: a hand-written agent loop plus four tools (read_file,
// write_file, edit_file, bash). Claude proposes tool calls; this program runs
// them, feeds the results back, and repeats until Claude stops asking.
//
//	export ANTHROPIC_API_KEY=sk-ant-...
//	go run .
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
)

const systemPrompt = `You are a coding assistant running inside a terminal harness.
Tools available: read_file, write_file, edit_file (exact unique string replace), and bash.
Use bash for searching (grep), listing, building, testing, and git.
Work in small steps. After editing code, build or test to verify your change. Keep edits minimal and explain briefly what you did.`

// shared reader so the REPL prompt and the bash confirmation never fight over stdin.
var stdin = bufio.NewReader(os.Stdin)

func main() {
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		fmt.Fprintln(os.Stderr, "error: set ANTHROPIC_API_KEY before running.")
		os.Exit(1)
	}
	model := os.Getenv("HARNESS_MODEL")
	if model == "" {
		model = "claude-opus-4-8"
	}

	client := anthropic.NewClient() // reads ANTHROPIC_API_KEY from env
	ctx := context.Background()
	tools := toolDefs()
	var messages []anthropic.MessageParam

	fmt.Printf("minimal-harness (model: %s) — type a request, Ctrl-D to quit.\n", model)
	for {
		fmt.Print("\n> ")
		line, err := stdin.ReadString('\n')
		if err != nil { // EOF
			fmt.Println()
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		messages = append(messages, anthropic.NewUserMessage(anthropic.NewTextBlock(line)))
		messages = runLoop(ctx, &client, model, tools, messages)
	}
}

// runLoop is the agentic loop: call the API, run any tool the model asks for,
// send the results back, and stop when the model is no longer asking for tools.
func runLoop(ctx context.Context, client *anthropic.Client, model string, tools []anthropic.ToolUnionParam, messages []anthropic.MessageParam) []anthropic.MessageParam {
	const maxSteps = 25 // guard: never loop (or bill) forever
	for step := 0; step < maxSteps; step++ {
		resp, err := client.Messages.New(ctx, anthropic.MessageNewParams{
			Model:     anthropic.Model(model),
			MaxTokens: 8192,
			System:    []anthropic.TextBlockParam{{Text: systemPrompt}},
			Messages:  messages,
			Tools:     tools,
		})
		if err != nil {
			fmt.Println("API error:", err)
			return messages
		}
		messages = append(messages, resp.ToParam()) // keep tool_use blocks in history

		var toolResults []anthropic.ContentBlockParamUnion
		for _, block := range resp.Content {
			switch b := block.AsAny().(type) {
			case anthropic.TextBlock:
				if strings.TrimSpace(b.Text) != "" {
					fmt.Println("\n" + b.Text)
				}
			case anthropic.ToolUseBlock:
				out, isErr := dispatch(b.Name, b.Input)
				fmt.Printf("  [%s] %s\n", b.Name, preview(out))
				toolResults = append(toolResults, anthropic.NewToolResultBlock(b.ID, out, isErr))
			}
		}

		if resp.StopReason != anthropic.StopReasonToolUse {
			return messages // model answered; done with this turn
		}
		messages = append(messages, anthropic.NewUserMessage(toolResults...))
	}
	fmt.Println("(stopped: reached max steps)")
	return messages
}

// dispatch runs one tool by name and returns (output, isError). Errors come back
// as tool results with isError=true so the model can adjust instead of crashing.
func dispatch(name string, input json.RawMessage) (string, bool) {
	switch name {
	case "read_file":
		var a struct {
			Path string `json:"path"`
		}
		json.Unmarshal(input, &a)
		data, err := os.ReadFile(a.Path)
		if err != nil {
			return err.Error(), true
		}
		return string(data), false

	case "write_file":
		var a struct {
			Path    string `json:"path"`
			Content string `json:"content"`
		}
		json.Unmarshal(input, &a)
		if err := os.WriteFile(a.Path, []byte(a.Content), 0o644); err != nil {
			return err.Error(), true
		}
		return "wrote " + a.Path, false

	case "edit_file":
		var a struct {
			Path string `json:"path"`
			Old  string `json:"old"`
			New  string `json:"new"`
		}
		json.Unmarshal(input, &a)
		data, err := os.ReadFile(a.Path)
		if err != nil {
			return err.Error(), true
		}
		content := string(data)
		switch strings.Count(content, a.Old) {
		case 0:
			return "old string not found in " + a.Path, true
		case 1:
			updated := strings.Replace(content, a.Old, a.New, 1)
			if err := os.WriteFile(a.Path, []byte(updated), 0o644); err != nil {
				return err.Error(), true
			}
			return "edited " + a.Path, false
		default:
			return "old string is not unique in " + a.Path + " (add more context to make it unique)", true
		}

	case "bash":
		var a struct {
			Command string `json:"command"`
		}
		json.Unmarshal(input, &a)
		if !confirm(a.Command) {
			return "user denied this command", true
		}
		out, err := exec.Command("bash", "-c", a.Command).CombinedOutput()
		if err != nil {
			return string(out) + "\n(exit error: " + err.Error() + ")", true
		}
		return string(out), false
	}
	return "unknown tool: " + name, true
}

// confirm gates a bash command behind a y/N prompt — the one human checkpoint.
func confirm(command string) bool {
	fmt.Printf("\n  run: %s\n  allow? [y/N] ", command)
	line, _ := stdin.ReadString('\n')
	return strings.EqualFold(strings.TrimSpace(line), "y")
}

func preview(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 200 {
		return s[:200] + "…"
	}
	if s == "" {
		return "(empty)"
	}
	return s
}

// toolDefs declares the four tools and their JSON-Schema inputs for the API.
func toolDefs() []anthropic.ToolUnionParam {
	str := map[string]any{"type": "string"}
	return []anthropic.ToolUnionParam{
		toolDef("read_file", "Read and return the full contents of a file.",
			map[string]any{"path": str}, "path"),
		toolDef("write_file", "Create a new file or overwrite an existing one with the given content.",
			map[string]any{"path": str, "content": str}, "path", "content"),
		toolDef("edit_file", "Replace an exact, unique substring in a file. 'old' must appear exactly once.",
			map[string]any{"path": str, "old": str, "new": str}, "path", "old", "new"),
		toolDef("bash", "Run a shell command (the user must confirm first). Use for grep/search, ls, build, test, and git.",
			map[string]any{"command": str}, "command"),
	}
}

func toolDef(name, desc string, props map[string]any, required ...string) anthropic.ToolUnionParam {
	return anthropic.ToolUnionParam{OfTool: &anthropic.ToolParam{
		Name:        name,
		Description: anthropic.String(desc),
		InputSchema: anthropic.ToolInputSchemaParam{
			Properties: props,
			Required:   required,
		},
	}}
}
