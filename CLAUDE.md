# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose
Learning resource for Claude Code mastery. Demonstrates skills, hooks, agents, commands, memory, and MCP configuration. No application code — only config and examples.

## Stack
- Platform: macOS (darwin)
- Shell: zsh
- Package manager: pnpm (prefer over npm/yarn)

## Structure
```
.claude/
  agents/       # Project-level agents (architect, security-reviewer)
  commands/     # Slash commands: /plan, /code-review, /security-scan
  settings.json # Hooks config (Stop + PostToolUse/Write → session.log)
  settings.local.json  # Auto-allow permissions, MCP enables
.mcp.json       # MCP servers: github, notion
examples/
  hooks/        # Example hook scripts (format-on-save, security-guard)
  skills/       # Example skill definitions
projects/
  zk-research/  # Sub-project with its own .claude config, Go code
CLAUDE_MASTERY.md  # Comprehensive reference guide for all Claude Code topics
```

## Sub-project: zk-research
Go module at `projects/zk-research/`. Has its own `.claude/` with agents (`zk-architect`, `zk-researcher`, `circuit-reviewer`), commands (`/research`, `/summarize`, `/circuit-review`), hooks, and skills.

```sh
cd projects/zk-research
go build ./...
go test ./...
go test ./tests/...          # Run specific test directory
go test -run TestMiMC ./...  # Run single test
```

## MCP Servers
- **github**: full GitHub API access (search, PR, issues, file contents)
- **notion**: requires real token in `.mcp.json` (currently placeholder)

## Rules
- Always read a file before editing it
- Prefer explicit over implicit code
- No emojis in output unless asked
- Lead with the answer, not the explanation
