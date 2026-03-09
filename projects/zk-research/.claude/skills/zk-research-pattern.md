# Skill: ZK Research Pattern

Standard workflow for researching a ZK primitive end-to-end.

## When to Use
Trigger when the user wants to learn about, implement, or evaluate a zero-knowledge primitive (e.g., "research Pedersen commitments", "understand Poseidon hash", "evaluate PlonK").

## Workflow

### Step 1: Literature Survey
Use the `zk-researcher` agent to:
- Define the primitive precisely (what relation it proves, what it computes)
- Identify the original paper and key follow-up work
- List cryptographic assumptions it relies on

### Step 2: Existing Notes Check
- Read all files in `notes/`
- Identify if this topic or a related one is already covered
- If yes, extend the existing note rather than creating a new one

### Step 3: gnark Mapping
- Identify which gnark types and functions are relevant
  - `frontend.API` methods: `api.Add`, `api.Mul`, `api.AssertIsEqual`, etc.
  - Relevant stdlib gadgets: `std/hash/mimc`, `std/algebra`, `std/commitments`
- Sketch the circuit `Define(api frontend.API)` method

### Step 4: Write the Note
- Create `notes/NN-<slug>.md` using the standard format from `zk-researcher`
- Include a "gnark Implementation Notes" section with a code sketch

### Step 5: Implement
- Write the circuit in `circuits/<category>/<name>.go`
- Write a test in `tests/<name>_test.go` that:
  1. Creates a valid witness and asserts proof succeeds
  2. Creates an invalid witness and asserts proof fails

### Step 6: Verify
```bash
go vet ./...
go test ./... -v
```

## Output Checklist
- [ ] Research note written in `notes/`
- [ ] Circuit implemented in `circuits/`
- [ ] Test passing: valid witness proves, invalid witness fails
- [ ] `go vet` passes
