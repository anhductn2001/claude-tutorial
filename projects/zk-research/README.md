# ZK Research Project

A research codebase for exploring zero-knowledge proof primitives using [gnark](https://github.com/consensys/gnark) (Go ZK framework by ConsenSys).

## Goals
- Implement canonical ZK circuits (MiMC hash, Merkle proofs, etc.)
- Generate and verify Groth16 proofs on BN254
- Build intuition for circuit design and constraint systems

## Stack
- **Language**: Go
- **ZK framework**: gnark v0.9+
- **Proving system**: Groth16 (BN254 curve)

## Project Layout

```
circuits/       gnark circuit definitions
scripts/        proof generation & verification scripts
tests/          Go test files
notes/          markdown research notes
.claude/        agents, commands, hooks, skills
```

## Quick Start

```bash
# Install dependencies
go mod download

# Run all tests
go test ./...

# Run MiMC circuit test
go test ./tests/ -run TestMiMC -v

# Generate and verify a proof
go run scripts/prove.go
```

## Claude Code Features Demonstrated
| Feature | Location | Purpose |
|---------|----------|---------|
| Agent | `.claude/agents/zk-researcher.md` | Research ZK topics |
| Agent | `.claude/agents/circuit-reviewer.md` | Audit circuit soundness |
| Agent | `.claude/agents/zk-architect.md` | Design ZK systems |
| Command | `.claude/commands/research.md` | `/research <topic>` |
| Command | `.claude/commands/circuit-review.md` | `/circuit-review` |
| Command | `.claude/commands/summarize.md` | `/summarize` |
| Hook | `.claude/hooks/lint-on-save.sh` | `go vet` on save |
| Skill | `.claude/skills/zk-research-pattern.md` | Research workflow |
