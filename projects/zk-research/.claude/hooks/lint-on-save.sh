#!/usr/bin/env bash
# Hook: PostToolUse/Write
# Runs `go vet` on the package containing any saved .go file.

set -euo pipefail

# CLAUDE_TOOL_INPUT_FILE_PATH is set by Claude Code for Write tool hooks
FILE="${CLAUDE_TOOL_INPUT_FILE_PATH:-}"

if [[ -z "$FILE" ]]; then
  exit 0
fi

# Only act on Go source files
if [[ "$FILE" != *.go ]]; then
  exit 0
fi

# Resolve the directory of the saved file
DIR="$(dirname "$FILE")"

echo "[lint-on-save] Running go vet on $DIR ..."

cd "$DIR"
go vet ./... 2>&1

echo "[lint-on-save] go vet passed."
