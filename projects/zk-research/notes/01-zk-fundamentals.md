# ZK Fundamentals

## Overview
Zero-knowledge proofs (ZKPs) allow a prover to convince a verifier that a statement is true without revealing any information beyond the truth of the statement. The three core properties are **completeness**, **soundness**, and **zero-knowledge**.

## Key Definitions

- **Statement**: A claim of the form "I know a witness w such that R(x, w) = 1" where x is public and w is private.
- **Prover**: Holds the witness w; generates the proof.
- **Verifier**: Holds only the public input x; checks the proof.
- **Relation R**: The NP relation defining what constitutes a valid witness.
- **Circuit**: An arithmetic representation of the relation R over a finite field.
- **R1CS**: Rank-1 Constraint System — the standard wire format for SNARKs. Each constraint has the form (a · w) * (b · w) = (c · w).
- **SNARK**: Succinct Non-interactive ARgument of Knowledge — proofs are short and verification is fast.

## Core Properties

### Completeness
If the prover holds a valid witness, they can always produce an accepted proof. Formally: for all valid (x, w), Pr[Verify(x, Prove(x,w)) = 1] = 1.

### Soundness
A cheating prover (without a valid witness) cannot produce an accepted proof, except with negligible probability (the *soundness error*).

### Zero-Knowledge
The verifier learns nothing about w beyond "a valid w exists." Formally: the verifier's view can be simulated without access to w.

## How It Works (High Level)

1. **Arithmetize**: Express the computation as polynomial equations over a finite field F_p.
2. **Commit**: The prover commits to the witness polynomial.
3. **Challenge**: The verifier (or random oracle) issues a challenge.
4. **Respond**: The prover evaluates the committed polynomial at the challenge point.
5. **Verify**: The verifier checks the evaluation against the committed value.

For SNARKs (e.g., Groth16), the interaction is compressed into a single non-interactive proof via the Fiat-Shamir heuristic or a trusted setup.

## Proving Systems Comparison

| System | Transparent | Proof Size | Verify Time | Notes |
|--------|------------|-----------|------------|-------|
| Groth16 | No (per-circuit trusted setup) | ~200 bytes | O(1) | Smallest proofs |
| PLONK | No (universal setup) | ~400 bytes | O(1) | Flexible |
| STARKs | Yes | ~100KB | O(log^2 n) | No trusted setup |

## gnark Implementation Notes

gnark represents circuits as Go structs implementing the `frontend.Circuit` interface:

```go
type MyCircuit struct {
    X frontend.Variable `gnark:",public"`  // public input
    W frontend.Variable                     // private witness
}

func (c *MyCircuit) Define(api frontend.API) error {
    // constraints go here
    api.AssertIsEqual(c.X, c.W)
    return nil
}
```

Key API methods:
- `api.Add(a, b, ...)` — field addition
- `api.Mul(a, b)` — field multiplication
- `api.AssertIsEqual(a, b)` — equality constraint
- `api.AssertIsBoolean(a)` — forces a ∈ {0, 1}

## Security Assumptions

- Groth16 relies on the **Knowledge of Exponent** assumption and the **discrete log** assumption in the target group.
- Soundness is in the **Random Oracle Model** (for non-interactive proofs).
- Trusted setup ceremonies produce a *Structured Reference String (SRS)*; if all participants collude, soundness breaks.

## References
- Groth 2016: "On the Size of Pairing-Based Non-interactive Arguments" — https://eprint.iacr.org/2016/260
- gnark documentation: https://docs.gnark.consensys.net
- ZKProof community standards: https://zkproof.org
