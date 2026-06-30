package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// ReadFile returns the contents of a file on the local filesystem.
type ReadFile struct{}

func (ReadFile) Name() string { return "read_file" }

func (ReadFile) Description() string {
	return "Read the contents of a file from the local filesystem and return it as text."
}

func (ReadFile) InputSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"path": map[string]any{
				"type":        "string",
				"description": "Path to the file to read.",
			},
		},
		"required": []string{"path"},
	}
}

func (ReadFile) Execute(_ context.Context, input json.RawMessage) (string, error) {
	var args struct {
		Path string `json:"path"`
	}
	if err := json.Unmarshal(input, &args); err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}
	if args.Path == "" {
		return "", fmt.Errorf("path is required")
	}
	data, err := os.ReadFile(args.Path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
