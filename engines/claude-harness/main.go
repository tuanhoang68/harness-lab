// Command claude-harness is a tiny agent harness built to learn how an agentic
// loop works: the model proposes tool calls, this program executes them, feeds
// the results back, and repeats until the model is done.
//
// Checkpoint 1 wires up the tools layer and verifies configuration. The agent
// loop (which talks to the Claude API) is added in the next checkpoint.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"claude-harness/tools"
)

func main() {
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		fmt.Fprintln(os.Stderr, "warning: ANTHROPIC_API_KEY is not set; export it before the agent loop is wired up.")
	}

	registry := tools.NewRegistry(
		tools.ReadFile{},
		tools.Bash{Confirm: confirmStdin},
	)

	fmt.Println("claude-harness — registered tools:")
	for _, t := range registry.All() {
		fmt.Printf("  - %-10s %s\n", t.Name(), t.Description())
	}
	fmt.Println("\nAgent loop not wired up yet (next checkpoint).")
}

// confirmStdin asks the user to approve a bash command on the terminal. It is
// the production Confirmer wired into the bash tool.
func confirmStdin(command string) bool {
	fmt.Printf("\nThe agent wants to run:\n  %s\nAllow? [y/N] ", command)
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.EqualFold(strings.TrimSpace(line), "y")
}
