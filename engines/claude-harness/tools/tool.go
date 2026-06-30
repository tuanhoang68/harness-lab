// Package tools defines the Tool interface the harness exposes to Claude, plus
// a small registry the agent loop uses to dispatch tool calls.
//
// Tools deliberately do NOT depend on the Anthropic SDK. A tool only knows how
// to describe itself (name, description, JSON-Schema for its input) and how to
// run. The agent layer translates these descriptions into SDK tool definitions
// and routes tool_use blocks back here. Keeping the boundary here means you can
// read, test, and reason about a tool without thinking about the wire protocol
// at all — and this package builds and tests with no network access.
package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
)

// Tool is one capability the agent can invoke. Claude sees Name/Description/
// InputSchema; the harness runs Execute when Claude asks for it.
type Tool interface {
	Name() string
	Description() string
	// InputSchema returns the JSON Schema (as a plain map) describing the tool's
	// input object. The agent layer hands this to the model verbatim.
	InputSchema() map[string]any
	// Execute runs the tool. input is the raw JSON the model produced for the
	// call; it conforms to InputSchema. The returned string is fed back to the
	// model as the tool_result. A non-nil error is surfaced to the model as an
	// error result so it can adapt rather than crashing the loop.
	Execute(ctx context.Context, input json.RawMessage) (string, error)
}

// Registry holds the tools available to an agent and dispatches calls by name.
type Registry struct {
	tools map[string]Tool
}

// NewRegistry builds a registry from the given tools. It panics on a duplicate
// name — that is a programming error, not a runtime condition.
func NewRegistry(ts ...Tool) *Registry {
	r := &Registry{tools: make(map[string]Tool, len(ts))}
	for _, t := range ts {
		if _, exists := r.tools[t.Name()]; exists {
			panic(fmt.Sprintf("tools: duplicate tool name %q", t.Name()))
		}
		r.tools[t.Name()] = t
	}
	return r
}

// All returns the registered tools in a stable (name-sorted) order, so the tool
// list handed to the model is deterministic — which keeps the prompt prefix
// byte-stable and lets caching work.
func (r *Registry) All() []Tool {
	names := make([]string, 0, len(r.tools))
	for name := range r.tools {
		names = append(names, name)
	}
	sort.Strings(names)

	out := make([]Tool, 0, len(names))
	for _, name := range names {
		out = append(out, r.tools[name])
	}
	return out
}

// Dispatch finds the named tool and runs it, returning (result, isError). An
// unknown tool name or an Execute error is reported back to the model as an
// error result rather than crashing the loop, so the model can recover.
func (r *Registry) Dispatch(ctx context.Context, name string, input json.RawMessage) (result string, isError bool) {
	t, ok := r.tools[name]
	if !ok {
		return fmt.Sprintf("error: unknown tool %q", name), true
	}
	out, err := t.Execute(ctx, input)
	if err != nil {
		return fmt.Sprintf("error: %v", err), true
	}
	return out, false
}
