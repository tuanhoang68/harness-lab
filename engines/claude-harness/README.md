# claude-harness

A tiny agent **harness** written in Go to learn how an agentic loop works from the
inside. The model (Claude) proposes tool calls; this program executes them, feeds
the results back, and repeats until the model is done. Claude decides *what* to do;
the harness is the part that *runs it, gates it, and stitches the results back* into
the conversation.

Built with the official Go SDK (`github.com/anthropics/anthropic-sdk-go`) but with a
**hand-written loop** — the loop is the thing worth learning, so it is not hidden
behind the SDK's tool runner.

## Concepts

- **Tool** — a capability the agent can invoke. Each tool describes itself (name,
  description, JSON-Schema input) and runs itself. Tools live in `tools/` and do
  **not** depend on the SDK, so they build and test offline.
- **Agent loop** (`agent/`, next checkpoint) — calls the API, checks `stop_reason`,
  dispatches `tool_use` blocks to the registry, sends `tool_result` back, and loops
  until the model stops asking for tools.
- **Gate** — the `bash` tool asks for confirmation before running a command. This is
  why an action gets promoted to a dedicated tool: so the harness can gate something
  hard to reverse, instead of letting the model run arbitrary shell unchecked.

## Tools

| Tool        | Input        | Does                                            |
| ----------- | ------------ | ----------------------------------------------- |
| `read_file` | `{path}`     | Returns the contents of a file.                 |
| `bash`      | `{command}`  | Runs a shell command (after y/N confirmation).  |

## Run

```sh
export ANTHROPIC_API_KEY=sk-ant-...
go run .
```

Checkpoint 1 just registers the tools and prints them. The interactive loop arrives
with the `agent` package.

## Build & test

```sh
go build ./...
go test ./...
```

The `tools/` package builds and tests with no network access.

## Status

See `docs/superpowers/specs/2026-06-18-claude-harness-learning-design.md` for the
design and the checkpoint plan. Current state: **checkpoint 1 — tools layer**.
