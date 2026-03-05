---
description: Review code for quality, bugs, security issues, and style consistency
argument-hint: optional file, PR number, or commit to review
---

Review the code changes since the last commit (or the file/PR: $ARGUMENTS).

Evaluate:
1. **Correctness**: Does the logic match the intent? Edge cases handled?
2. **Security**: OWASP Top 10 issues? Input validated? No hardcoded secrets?
3. **Performance**: N+1 queries? Blocking I/O? Memory leaks?
4. **Tests**: Adequate coverage? Are failures handled?
5. **Style**: Consistent with project conventions in CLAUDE.md?

Output as:
## BLOCKER
- [file:line] Issue

## SUGGESTION
- [file:line] Improvement

## NIT
- [file:line] Minor issue
