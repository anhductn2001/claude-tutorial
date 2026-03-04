#!/bin/bash
# PreToolUse hook: block dangerous Bash commands
# Configure in .claude/settings.json under "PreToolUse" with matcher "Bash"
#
# Hook exit codes:
#   0 = allow
#   1 = block (Claude sees the stderr as error)
#   2 = block silently

COMMAND="$CLAUDE_TOOL_INPUT"

# Block destructive operations that should never run automatically
DANGEROUS_PATTERNS=(
  "rm -rf /"
  "rm -rf ~"
  "git push --force.*main"
  "git push --force.*master"
  "DROP TABLE"
  "DROP DATABASE"
  "truncate.*--yes"
  "kubectl delete"
)

for pattern in "${DANGEROUS_PATTERNS[@]}"; do
  if echo "$COMMAND" | grep -qiE "$pattern"; then
    echo "BLOCKED by security-guard: Pattern '$pattern' matched" >&2
    echo "Command: $COMMAND" >&2
    exit 1
  fi
done

exit 0
