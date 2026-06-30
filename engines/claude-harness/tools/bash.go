package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

// Confirmer decides whether a bash command may run. Returning false blocks it.
// Injecting this makes the gate testable and lets main wire it to the terminal.
type Confirmer func(command string) bool

// Bash runs shell commands. Because a shell command is hard to reverse and can
// do almost anything, the harness gates every call behind a Confirmer. This is
// the "promote an action to a dedicated tool so the harness can gate it" lesson
// made concrete: the model emits a command, but the harness decides whether it
// runs.
type Bash struct {
	// Confirm is asked before each command runs. If nil, commands run without a
	// prompt (handy in tests; main always supplies one).
	Confirm Confirmer
}

func (Bash) Name() string { return "bash" }

func (Bash) Description() string {
	return "Run a bash command and return its combined stdout and stderr. " +
		"Use it to inspect the system, run builds or tests, and perform file operations."
}

func (Bash) InputSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"command": map[string]any{
				"type":        "string",
				"description": "The bash command to execute.",
			},
		},
		"required": []string{"command"},
	}
}

func (b Bash) Execute(ctx context.Context, input json.RawMessage) (string, error) {
	var args struct {
		Command string `json:"command"`
	}
	if err := json.Unmarshal(input, &args); err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}
	if args.Command == "" {
		return "", fmt.Errorf("command is required")
	}

	// The gate. A denial is a normal result, not an error — the model is told it
	// was rejected so it can choose a different approach.
	if b.Confirm != nil && !b.Confirm(args.Command) {
		return "command rejected by user", nil
	}

	out, err := exec.CommandContext(ctx, "bash", "-c", args.Command).CombinedOutput()
	if err != nil {
		// A non-zero exit is information for the model, not a harness failure:
		// return the output plus the failure so it can read stderr and adapt.
		return fmt.Sprintf("%s\n(command failed: %v)", out, err), nil
	}
	return string(out), nil
}
