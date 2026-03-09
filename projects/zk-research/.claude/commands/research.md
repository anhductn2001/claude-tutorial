# /research

Deep-dive into a ZK topic and write a structured research note.

## Usage
```
/research <topic>
```

## What This Does
1. Uses the `zk-researcher` agent to research `$ARGUMENTS`
2. Checks `notes/` for existing related notes to avoid duplication
3. Writes a new note file in `notes/` using the standard format
4. Summarizes key takeaways at the end

## Steps
1. Read all existing files in `notes/` to understand what has already been covered
2. Invoke the `zk-researcher` agent with the topic: `$ARGUMENTS`
3. Determine the note filename: `notes/NN-<slug>.md` where NN is the next sequential number
4. Write the note using the format defined in `zk-researcher` agent
5. Output a 3-bullet summary of the most important insights

## Example
```
/research MiMC hash function
/research Groth16 proving system
/research R1CS constraint system
```
