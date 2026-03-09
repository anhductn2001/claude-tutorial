# ZK Research Project — Claude Rules

## Stack
- Language: Go
- ZK Framework: gnark (github.com/consensys/gnark)
- Proving system: Groth16 on BN254 curve
- Test: Go standard testing package

## Rules
- Always read a file before editing it
- Run `go test ./...` after any code change
- Run `go vet ./...` before committing
- Prefer explicit over implicit code; no magic
- No emojis in output unless asked

## Circuit Safety
- Every signal in a gnark circuit MUST be constrained
- Use `api.AssertIsEqual` or boolean constraints — never leave signals unconstrained
- Document why each constraint is necessary

## Agents
- Use `zk-researcher` for literature/concept research
- Use `circuit-reviewer` to audit circuit soundness
- Use `zk-architect` for system-level design decisions

## Commands
- `/research <topic>` — deep dive into a ZK topic and write a note
- `/circuit-review` — audit all circuits in `circuits/`
- `/summarize` — synthesize all notes in `notes/`

## Directory Layout
- `circuits/` — gnark circuit definitions
- `scripts/` — prove/verify scripts
- `tests/` — Go test files
- `notes/` — markdown research notes
- `.claude/` — agents, commands, hooks, skills
