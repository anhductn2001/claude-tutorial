#!/bin/bash
# PostToolUse hook: auto-format files after Claude writes them
# Configure in .claude/settings.json under "PostToolUse" with matcher "Write"

FILE="$1"  # File path from hook environment
if [ -z "$FILE" ]; then
  exit 0
fi

case "$FILE" in
  *.ts|*.tsx|*.js|*.jsx)
    command -v prettier &>/dev/null && prettier --write "$FILE"
    ;;
  *.py)
    command -v black &>/dev/null && black "$FILE"
    ;;
  *.go)
    gofmt -w "$FILE"
    ;;
  *.rs)
    command -v rustfmt &>/dev/null && rustfmt "$FILE"
    ;;
esac
