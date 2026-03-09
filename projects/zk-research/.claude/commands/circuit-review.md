# /circuit-review

Audit all gnark circuits in the `circuits/` directory for security issues.

## Usage
```
/circuit-review
/circuit-review circuits/hash/mimc.go
```

## What This Does
1. Discovers all `.go` files under `circuits/`
2. Uses the `circuit-reviewer` agent to audit each file
3. Produces a consolidated security report

## Steps
1. List all `.go` files in `circuits/` recursively
2. For each file, invoke the `circuit-reviewer` agent to audit it
3. Aggregate findings by severity: CRITICAL, WARNING, INFO
4. Output a report with:
   - Summary table (file → finding count by severity)
   - Detailed findings per file
   - Recommended fixes for CRITICAL issues

## Output Format
```
## Circuit Security Report

### Summary
| File | Critical | Warning | Info |
|------|---------|---------|------|
| circuits/hash/mimc.go | 0 | 1 | 2 |

### Detailed Findings
...
```

## Notes
- CRITICAL findings must be fixed before any proving
- Run `go vet ./...` and `go test ./...` after applying fixes
