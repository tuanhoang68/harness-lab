package tools

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "hello.txt")
	want := "hello harness\n"
	if err := os.WriteFile(path, []byte(want), 0o644); err != nil {
		t.Fatal(err)
	}

	input, _ := json.Marshal(map[string]string{"path": path})
	got, err := ReadFile{}.Execute(context.Background(), input)
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReadFileMissing(t *testing.T) {
	input, _ := json.Marshal(map[string]string{"path": "/no/such/file"})
	if _, err := (ReadFile{}).Execute(context.Background(), input); err == nil {
		t.Fatal("expected an error for a missing file, got nil")
	}
}

func TestBashConfirmAllows(t *testing.T) {
	b := Bash{Confirm: func(string) bool { return true }}
	input, _ := json.Marshal(map[string]string{"command": "echo hi"})
	got, err := b.Execute(context.Background(), input)
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}
	if !strings.Contains(got, "hi") {
		t.Errorf("got %q, want output containing %q", got, "hi")
	}
}

func TestBashConfirmDenies(t *testing.T) {
	b := Bash{Confirm: func(string) bool { return false }}
	input, _ := json.Marshal(map[string]string{"command": "echo should-not-run"})
	got, err := b.Execute(context.Background(), input)
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}
	if strings.Contains(got, "should-not-run") {
		t.Errorf("command ran despite denial; got %q", got)
	}
}

func TestRegistryDispatchUnknown(t *testing.T) {
	r := NewRegistry(ReadFile{})
	got, isErr := r.Dispatch(context.Background(), "does_not_exist", nil)
	if !isErr {
		t.Error("expected isError=true for an unknown tool")
	}
	if !strings.Contains(got, "unknown tool") {
		t.Errorf("got %q, want an unknown-tool message", got)
	}
}

func TestRegistryAllSorted(t *testing.T) {
	r := NewRegistry(Bash{}, ReadFile{})
	all := r.All()
	if len(all) != 2 {
		t.Fatalf("got %d tools, want 2", len(all))
	}
	// All() must return a deterministic, name-sorted order.
	if all[0].Name() != "bash" || all[1].Name() != "read_file" {
		t.Errorf("got order [%s %s], want [bash read_file]", all[0].Name(), all[1].Name())
	}
}

func TestNewRegistryPanicsOnDuplicate(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected a panic on duplicate tool names")
		}
	}()
	NewRegistry(ReadFile{}, ReadFile{})
}
