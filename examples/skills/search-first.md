# Skill: search-first (Research-First Development)

## Trigger
TRIGGER when: user asks to implement a feature, fix a bug, or refactor existing code.
DO NOT TRIGGER when: the request is purely conversational or asks for an explanation only.

## Instructions

### Phase 1: Research (ALWAYS do this first)
Before writing a single line of code:

1. Read CLAUDE.md for project conventions and stack details
2. Search codebase for related existing code:
   - Grep for the function/class/pattern name
   - Look for similar implementations to follow as a model
3. Read the 2-3 most relevant files in full
4. Check git log for context on recent changes to relevant files
5. Identify: What already exists? What needs to change? What are the risks?

### Phase 2: Report Findings
Summarize to the user:
- What exists today (existing code, patterns, relevant files)
- What needs to change (specific files and functions)
- Proposed approach (concrete, not vague)
- Unknowns or risks (flag these explicitly)

### Phase 3: Get Confirmation
Ask: "Does this approach look right before I start coding?"

### Phase 4: Implement
Only after confirmation. Reference the patterns found in Phase 1.
Follow existing conventions exactly — don't introduce new patterns without asking.

## Anti-Patterns to Avoid
- Writing code before reading the relevant files
- Assuming a library/API works a certain way without verifying
- Introducing new patterns when existing ones exist
- Making changes to files that weren't read first
