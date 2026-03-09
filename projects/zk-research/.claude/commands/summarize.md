# /summarize

Synthesize all research notes in `notes/` into a coherent overview.

## Usage
```
/summarize
```

## What This Does
1. Reads all markdown files in `notes/`
2. Identifies themes, connections, and open questions
3. Produces a synthesis document

## Steps
1. List and read all files in `notes/`
2. Use the `zk-researcher` agent to identify:
   - Common themes across notes
   - Connections between concepts
   - Open questions that need more research
   - Recommended next topics to study
3. Output a structured synthesis (do NOT write to a file — just display)

## Output Format
```markdown
## ZK Research Synthesis

### Topics Covered
- Topic A (notes/01-...) — one-line summary
- Topic B (notes/02-...) — one-line summary

### Key Connections
- How A relates to B
- ...

### Open Questions
- Question 1
- Question 2

### Recommended Next Topics
1. Topic X — why it matters
2. Topic Y — why it matters
```
