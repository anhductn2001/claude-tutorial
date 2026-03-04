# Claude Code Mastery Guide

A practical, production-ready reference covering every major Claude Code concept.
Based on: https://github.com/affaan-m/everything-claude-code

---

## Table of Contents
1. [Project Structure](#1-project-structure)
2. [CLAUDE.md — Rules & Instructions](#2-claudemd--rules--instructions)
3. [Skills](#3-skills)
4. [Slash Commands](#4-slash-commands)
5. [Hooks](#5-hooks)
6. [Agents (Subagents)](#6-agents-subagents)
7. [Memory Optimization](#7-memory-optimization)
8. [Continuous Learning & Instincts](#8-continuous-learning--instincts)
9. [Security Scanning](#9-security-scanning)
10. [Research-First Development](#10-research-first-development)
11. [MCP Configuration](#11-mcp-configuration)
12. [Token Optimization](#12-token-optimization)

---

## 1. Project Structure

Claude Code looks for configuration in two places:

```
~/.claude/                     # Global (user-level) config
├── CLAUDE.md                  # Global rules always sent to Claude
├── settings.json              # Global settings
├── commands/                  # Global slash commands
├── agents/                    # Global subagent definitions
└── hooks/                     # Global event hooks

your-project/
├── CLAUDE.md                  # Project-level rules (merged with global)
├── .claude/
│   ├── settings.json          # Project settings (overrides global)
│   ├── commands/              # Project slash commands
│   ├── agents/                # Project agents
│   └── hooks/                 # Project hooks
```

**Load order**: Global CLAUDE.md → Project CLAUDE.md → Conversation

---

## 2. CLAUDE.md — Rules & Instructions

`CLAUDE.md` is always injected into every conversation. It's your "always follow" ruleset.

### Global rules (~/.claude/CLAUDE.md)
```markdown
# Global Rules

## Code Style
- Use 2-space indentation for JS/TS, 4-space for Python/Go
- Prefer explicit over implicit
- No magic numbers — use named constants

## Git Workflow
- Always write descriptive commit messages
- Never force-push to main/master
- Run tests before committing

## Security
- Never hardcode secrets
- Validate all external inputs
- Use parameterized queries for DB operations

## Communication
- Be concise — lead with the answer
- No emojis unless asked
- Prefer code examples over prose explanations
```

### Project rules (your-project/CLAUDE.md)
```markdown
# Project: my-api

## Stack
- Runtime: Node.js 20 + TypeScript
- Framework: Fastify
- DB: PostgreSQL with Drizzle ORM
- Tests: Vitest

## Conventions
- All API handlers in src/handlers/
- Business logic in src/services/
- Always return typed responses using zod schemas
- Use pnpm (not npm or yarn)

## Commands
- Build: `pnpm build`
- Test: `pnpm test`
- Lint: `pnpm lint`
```

---

## 3. Skills

Skills are reusable, parameterized workflow prompts stored as `.md` files.
They are invoked via the `Skill` tool by Claude itself (not by users directly).

### File location
```
~/.claude/skills/         # Global skills
.claude/skills/           # Project skills
```

### Anatomy of a skill

```markdown
# Skill: tdd-workflow

## Trigger
TRIGGER when: user asks to implement a feature with tests, or says "TDD"

## Workflow
1. Write failing test first
2. Implement minimum code to pass
3. Refactor with green tests
4. Repeat

## Instructions
- Use the test framework specified in CLAUDE.md
- Run tests after each step: `{test_command}`
- Never write implementation before tests exist
```

### Example: security-checklist skill
```markdown
# Skill: security-checklist

## Trigger
TRIGGER when: code imports auth libraries, handles user input, or user says "secure"

## Checklist
- [ ] No hardcoded credentials
- [ ] Input validated at boundaries
- [ ] SQL uses parameterized queries
- [ ] JWT secrets from environment variables
- [ ] Rate limiting on auth endpoints
- [ ] CORS configured restrictively
```

### Example: research-first skill
```markdown
# Skill: search-first

## Trigger
TRIGGER when: user asks to implement something complex or unfamiliar

## Instructions
1. BEFORE writing code, gather context:
   - Search codebase for existing patterns
   - Read relevant existing files
   - Understand the data model
2. Summarize findings to user
3. Propose approach and get confirmation
4. Then implement
```

---

## 4. Slash Commands

Commands are markdown files that expand into full prompts when invoked with `/command-name`.

### File location
```
~/.claude/commands/         # Global commands
.claude/commands/           # Project commands
```

### Example commands

**.claude/commands/plan.md**
```markdown
Analyze the current task and create a detailed implementation plan:

1. Identify all files that need to change
2. List dependencies and potential blockers
3. Break into atomic steps
4. Estimate complexity (S/M/L)
5. Highlight risks

Present as a numbered checklist. Ask for approval before starting.
```

**.claude/commands/code-review.md**
```markdown
Perform a thorough code review of the changes since the last commit:

1. Correctness: Does the logic do what's intended?
2. Security: Any OWASP Top 10 issues?
3. Performance: Any N+1 queries, blocking I/O, memory leaks?
4. Tests: Adequate coverage? Edge cases handled?
5. Style: Consistent with project conventions?

Output as GitHub-style review comments grouped by severity: BLOCKER, SUGGESTION, NIT.
```

**.claude/commands/security-scan.md**
```markdown
Run a security audit on the current codebase:

Check for:
- Hardcoded secrets, API keys, passwords
- SQL injection vulnerabilities
- XSS attack vectors
- Insecure dependencies (check package.json)
- Exposed sensitive endpoints
- Missing authentication/authorization
- Unsafe deserialization

Report findings grouped by: CRITICAL, HIGH, MEDIUM, LOW
```

**.claude/commands/tdd.md**
```markdown
Guide me through Test-Driven Development for: $ARGUMENTS

Steps:
1. Write a failing test that defines expected behavior
2. Run tests to confirm they fail (red)
3. Write minimum code to make tests pass (green)
4. Refactor while keeping tests green
5. Repeat

Always run tests between each step.
```

**.claude/commands/build-fix.md**
```markdown
The build is broken. Fix it:

1. Run the build command and capture full error output
2. Identify root cause (not just symptoms)
3. Fix the underlying issue
4. Verify build passes
5. Check that tests still pass

Do NOT use --force, --no-verify, or skip any safety checks.
```

---

## 5. Hooks

Hooks are shell scripts triggered by Claude Code lifecycle events.

### File location
```
~/.claude/hooks/                    # Global hooks
.claude/hooks/                      # Project hooks
```

### Hook configuration (.claude/settings.json)

```json
{
  "hooks": {
    "SessionStart": [
      {
        "matcher": "",
        "hooks": [
          {
            "type": "command",
            "command": "cat .claude/context.md 2>/dev/null || echo 'No saved context'"
          }
        ]
      }
    ],
    "PreToolUse": [
      {
        "matcher": "Bash",
        "hooks": [
          {
            "type": "command",
            "command": "echo 'Running: $CLAUDE_TOOL_INPUT'"
          }
        ]
      }
    ],
    "PostToolUse": [
      {
        "matcher": "Write",
        "hooks": [
          {
            "type": "command",
            "command": ".claude/hooks/format-on-save.sh"
          }
        ]
      }
    ],
    "Stop": [
      {
        "matcher": "",
        "hooks": [
          {
            "type": "command",
            "command": ".claude/hooks/save-context.sh"
          }
        ]
      }
    ]
  }
}
```

### Available hook events
| Event | Trigger |
|-------|---------|
| `SessionStart` | Claude Code session begins |
| `PreToolUse` | Before any tool call |
| `PostToolUse` | After any tool call |
| `Stop` | Claude stops responding |
| `SubagentStop` | Subagent finishes |
| `PreCompact` | Before context compression |

### Example hook scripts

**.claude/hooks/save-context.sh** — Persist session state
```bash
#!/bin/bash
# Called on Stop — save a summary for next session
cat > .claude/context.md << 'EOF'
# Last Session Context
Date: $(date)
Branch: $(git branch --show-current)
Last commit: $(git log -1 --oneline)
EOF
```

**.claude/hooks/format-on-save.sh** — Auto-format after Write
```bash
#!/bin/bash
# Format the file that was just written
FILE="$CLAUDE_TOOL_RESULT_PATH"
case "$FILE" in
  *.ts|*.tsx) npx prettier --write "$FILE" ;;
  *.py) python -m black "$FILE" ;;
  *.go) gofmt -w "$FILE" ;;
esac
```

**.claude/hooks/security-guard.sh** — Block dangerous patterns pre-tool
```bash
#!/bin/bash
# PreToolUse hook for Bash — block dangerous commands
INPUT="$CLAUDE_TOOL_INPUT"
if echo "$INPUT" | grep -qE "rm -rf|DROP TABLE|git push --force"; then
  echo "BLOCKED: Dangerous command detected" >&2
  exit 1
fi
```

---

## 6. Agents (Subagents)

Agents are specialized Claude instances with a focused system prompt.
They run in isolation and return results to the parent conversation.

### File location
```
~/.claude/agents/          # Global agents
.claude/agents/            # Project agents
```

### Agent definition anatomy

**.claude/agents/code-reviewer.md**
```markdown
---
name: code-reviewer
description: Reviews code for quality, bugs, and security issues. Use when a PR or feature is complete.
---

You are an expert code reviewer. Your job is to:

1. Find bugs, logic errors, and edge cases
2. Identify security vulnerabilities (OWASP Top 10)
3. Check for performance issues (N+1, memory leaks)
4. Verify test coverage
5. Ensure consistency with project style

Output format:
## Blockers
- [file:line] Issue description

## Suggestions
- [file:line] Improvement description

## Nitpicks
- [file:line] Minor style issue

Be direct and specific. Reference line numbers.
```

**.claude/agents/security-reviewer.md**
```markdown
---
name: security-reviewer
description: Deep security audit. Use after code-reviewer for auth, crypto, or data handling code.
---

You are a senior security engineer. Audit this code for:

- Authentication bypass vulnerabilities
- Authorization flaws (IDOR, privilege escalation)
- Injection attacks (SQL, NoSQL, command, LDAP)
- Sensitive data exposure
- Broken cryptography
- Security misconfigurations
- Insecure direct object references

For each finding, provide:
- Severity: CRITICAL/HIGH/MEDIUM/LOW
- CWE reference
- Proof-of-concept attack scenario
- Recommended fix
```

**.claude/agents/architect.md**
```markdown
---
name: architect
description: System design and architecture decisions. Use for new features or refactoring.
---

You are a principal software architect. When designing systems:

1. Start with requirements and constraints
2. Consider scalability, reliability, maintainability
3. Evaluate 2-3 alternatives with trade-offs
4. Recommend the simplest solution that meets requirements
5. Identify future inflection points

Output as:
- Decision: What to build
- Rationale: Why this approach
- Trade-offs: What you're giving up
- Risks: What could go wrong
```

### Using agents in practice
Claude Code automatically invokes agents when their description matches.
The `description` field is used for routing — make it specific.

---

## 7. Memory Optimization

### The auto-memory system
Claude Code maintains a `memory/` directory in your project's context:
```
~/.claude/projects/<project-hash>/memory/
├── MEMORY.md       # Always loaded (keep under 200 lines)
├── patterns.md     # Recurring code patterns
├── debugging.md    # Solved issues and root causes
└── decisions.md    # Architecture decisions made
```

### MEMORY.md best practices
```markdown
# Project Memory

## Key Files
- Entry point: src/main.ts
- Config: src/config/index.ts
- DB models: src/db/schema.ts

## Conventions Confirmed
- Use Drizzle ORM (not Prisma) — migration path blocked
- Error handling uses Result<T, E> pattern
- All API responses typed with Zod schemas

## Recurring Patterns
- Auth middleware: src/middleware/auth.ts:checkJWT()
- Pagination: always use cursor-based (not offset)

## Known Issues
- tests/e2e/ skip in CI (needs Docker) — use unit tests only
- Package manager: pnpm (bun breaks some scripts)
```

### Context compaction strategy
- Use `/clear` between unrelated tasks
- Compact at logical breakpoints (after finishing a feature)
- Set `CLAUDE_AUTOCOMPACT_PCT_OVERRIDE=50` for aggressive compaction
- Pre-compact hook saves state before compression

### settings.json for memory
```json
{
  "model": "sonnet",
  "MAX_THINKING_TOKENS": "10000",
  "CLAUDE_AUTOCOMPACT_PCT_OVERRIDE": "50"
}
```

---

## 8. Continuous Learning & Instincts

The instinct system extracts recurring patterns from sessions and evolves them into skills.

### How it works
1. **Collection**: During sessions, note recurring patterns ("I always do X before Y")
2. **Extraction**: `/instinct-status` shows collected instincts with confidence scores
3. **Evolution**: `/evolve` clusters instincts into reusable skills
4. **Import/Export**: Share instincts across projects

### Manual instinct capture
In your MEMORY.md or patterns.md:
```markdown
## Instincts (High Confidence)
- Always read the file before editing it (learned from 5+ failed edits)
- Run tests after every change, not just at the end
- Check package.json before suggesting a new dependency

## Instincts (Medium Confidence)
- TypeScript generics over `any` — project team confirmed preference
- Prefer `interface` over `type` for object shapes in this project
```

### Evolving instincts into skills
Once a pattern is confirmed across 3+ sessions, codify it:
```markdown
# Skill: pre-edit-check

## Trigger
TRIGGER when: about to edit any existing file

## Instructions
1. Always Read the file first
2. Check if similar patterns exist nearby
3. Only then make the targeted change
4. Verify no regressions with a quick grep
```

---

## 9. Security Scanning

### AgentShield approach
Integrate security as a first-class concern, not an afterthought.

### Security-first CLAUDE.md rules
```markdown
## Security Rules (Always Follow)

1. **No secrets in code**: Never commit API keys, passwords, tokens
   - Use environment variables
   - Add .env to .gitignore immediately

2. **Input validation**: Validate at ALL entry points
   - User forms, API endpoints, file uploads, CLI args
   - Use Zod/Joi/Pydantic at the boundary

3. **SQL safety**: Always use parameterized queries
   - Never concatenate user input into SQL strings

4. **Auth checks**: Every protected route needs explicit authz check
   - Don't rely on "the UI won't show that button"

5. **Dependency hygiene**: Check for known vulns before adding packages
   - Run `npm audit` / `pip audit` / `go vuln` after adding deps
```

### Security scan command (.claude/commands/security-scan.md)
```markdown
Perform a security audit. Check for:

**Secrets & Credentials**
- Hardcoded API keys, passwords, tokens, private keys
- .env files accidentally committed
- Credentials in logs or error messages

**Injection**
- SQL injection (string concatenation in queries)
- Command injection (unsanitized shell exec)
- XSS (unescaped user content in HTML)

**Authentication & Authorization**
- Missing auth checks on endpoints
- IDOR vulnerabilities (can user A access user B's data?)
- Weak JWT configuration

**Dependencies**
- Run: `npm audit --json` or equivalent
- Flag CRITICAL and HIGH severity

Output as: CRITICAL / HIGH / MEDIUM / LOW with file:line references.
```

### Security agent
```markdown
---
name: security-reviewer
description: Run after implementing auth, user input handling, file uploads, or payment flows
---

Audit the provided code specifically for:
1. Authentication bypass
2. Authorization escalation
3. Data injection (SQL, NoSQL, command)
4. Sensitive data exposure in logs/responses
5. Insecure cryptography choices

Give CVSS severity scores and remediation steps.
```

---

## 10. Research-First Development

The core principle: **investigate before you implement**.

### The search-first workflow

1. **Understand the codebase** before touching it
   ```
   - Grep for the relevant function/module
   - Read the files involved
   - Check git log for context
   ```

2. **Validate assumptions** before coding
   ```
   - Does this library/API work the way I think?
   - Does a similar pattern already exist?
   - What are the edge cases?
   ```

3. **Propose, then implement**
   ```
   - Summarize findings
   - Propose approach
   - Get confirmation
   - Then write code
   ```

### Research-first skill (.claude/skills/search-first.md)
```markdown
# Skill: search-first

## Trigger
TRIGGER when: user asks to implement, fix, or refactor anything non-trivial

## Instructions

### Phase 1: Research (Do this FIRST)
1. Read CLAUDE.md for project conventions
2. Search codebase for related patterns: Grep for relevant terms
3. Read the most relevant 2-3 files fully
4. Check if the feature/fix already partially exists

### Phase 2: Report
Summarize:
- What exists already
- What needs to change
- Proposed approach
- Risks or unknowns

### Phase 3: Implement
Only after user confirms the approach.
```

---

## 11. MCP Configuration

MCP (Model Context Protocol) servers extend Claude with external tools.

### Configuration location
```json
// ~/.claude/settings.json (global) or .claude/settings.json (project)
{
  "mcpServers": {
    "github": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-github"],
      "env": {
        "GITHUB_TOKEN": "${GITHUB_TOKEN}"
      }
    },
    "filesystem": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-filesystem", "/path/to/project"]
    },
    "postgres": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-postgres"],
      "env": {
        "DATABASE_URL": "${DATABASE_URL}"
      }
    }
  }
}
```

### Common MCP servers
| Server | Purpose | Install |
|--------|---------|---------|
| `@modelcontextprotocol/server-github` | GitHub API | npx |
| `@modelcontextprotocol/server-filesystem` | File system | npx |
| `@modelcontextprotocol/server-postgres` | PostgreSQL | npx |
| `@modelcontextprotocol/server-sqlite` | SQLite | npx |
| `@modelcontextprotocol/server-brave-search` | Web search | npx |

### MCP security rules
- Never put real credentials in settings.json (use env vars)
- Limit MCP servers to <10 per project (token overhead)
- Only enable MCPs relevant to the current task
- Review MCP server code before using unknown servers

---

## 12. Token Optimization

### Model selection
```json
{
  "model": "sonnet"
}
```
- **Sonnet** (default): Best for most coding tasks, ~60% cheaper than Opus
- **Opus**: Complex reasoning, architecture decisions
- **Haiku**: Simple, fast tasks — searching, lookups, formatting

### Compaction strategy
```json
{
  "CLAUDE_AUTOCOMPACT_PCT_OVERRIDE": "50"
}
```
- Compact at 50% context = more frequent but cheaper
- Use `/clear` between unrelated tasks
- Pre-compact hook: save state before compression

### Minimize context bloat
1. Keep MEMORY.md under 200 lines
2. Use focused agents (isolated context windows)
3. Don't load large files unless needed
4. Limit active MCP servers

### Parallel agents = faster + cheaper
For independent tasks, use parallel subagents:
```
/multi-execute
- Agent 1: Review auth module
- Agent 2: Review payment module
- Agent 3: Review API endpoints
```
Each agent has its own context window — no cross-contamination.

---

## Quick Reference: File Structure to Create

```
~/.claude/
├── CLAUDE.md                    # Your global always-follow rules
├── settings.json                # Global settings + MCPs
├── commands/
│   ├── plan.md                  # /plan command
│   ├── code-review.md           # /code-review command
│   ├── security-scan.md         # /security-scan command
│   ├── tdd.md                   # /tdd command
│   └── build-fix.md             # /build-fix command
└── agents/
    ├── code-reviewer.md         # Code quality agent
    ├── security-reviewer.md     # Security audit agent
    └── architect.md             # Architecture decisions agent

your-project/.claude/
├── settings.json                # Project settings + hooks
├── commands/
│   └── deploy.md                # Project-specific commands
└── agents/
    └── db-reviewer.md           # Project-specific agents

your-project/CLAUDE.md           # Project rules + stack info
```

---

## Next Steps

1. Create `~/.claude/CLAUDE.md` with your global rules
2. Add a project `CLAUDE.md` with stack details
3. Create 3-5 commands you use repeatedly
4. Set up a `Stop` hook to save session context
5. Define a `security-reviewer` agent for sensitive code
6. Establish `MEMORY.md` patterns for your project
