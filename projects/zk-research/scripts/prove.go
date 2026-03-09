// Command prove generates and verifies a Groth16 proof for the MiMC circuit.
package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"

	"github.com/decentrio/zk-research/circuits/hash"
)

func main() {
	// 1. Compile the circuit into an R1CS constraint system
	fmt.Println("Compiling MiMC circuit...")
	var circuit hash.MiMCCircuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		log.Fatalf("compile error: %v", err)
	}
	fmt.Printf("Constraints: %d\n", ccs.GetNbConstraints())

	// 2. Groth16 trusted setup (in production: use a real ceremony SRS)
	fmt.Println("Running trusted setup (Groth16)...")
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		log.Fatalf("setup error: %v", err)
	}

	// 3. Build a valid witness
	// We use a simple preimage = 42, key = 0 for demo purposes.
	// The expected hash must be computed off-circuit; here we cheat by
	// running the circuit in "test" mode to extract the hash value.
	//
	// In a real application you would compute the hash natively in Go
	// using gnark-crypto's MiMC implementation.
	preimage := big.NewInt(42)
	key := big.NewInt(0)

	// Compute expected hash natively using big.Int field arithmetic.
	expectedHash := computeMiMCNative(preimage, key)
	fmt.Printf("Preimage: %s, Key: %s, Hash: %s\n", preimage, key, expectedHash)

	assignment := hash.MiMCCircuit{
		Preimage: preimage,
		Key:      key,
		Hash:     expectedHash,
	}

	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		log.Fatalf("witness error: %v", err)
	}

	publicWitness, err := witness.Public()
	if err != nil {
		log.Fatalf("public witness error: %v", err)
	}

	// 4. Generate proof
	fmt.Println("Generating proof...")
	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		log.Fatalf("prove error: %v", err)
	}
	fmt.Println("Proof generated.")

	// 5. Verify proof
	fmt.Println("Verifying proof...")
	if err := groth16.Verify(proof, vk, publicWitness); err != nil {
		log.Fatalf("verify error: %v", err)
	}
	fmt.Println("Proof verified successfully.")
}

// bn254ScalarField is the BN254 scalar field modulus.
var bn254ScalarField, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088548364400416034343698204186575808495617", 10,
)

// computeMiMCNative mirrors the circuit's 7-round MiMC using big.Int field arithmetic.
// This is a simplified demo circuit — not production-safe.
func computeMiMCNative(preimage, key *big.Int) *big.Int {
	p := bn254ScalarField
	result := new(big.Int).Set(preimage)
	for c := 0; c < 7; c++ {
		// t = (result + key + c) mod p
		t := new(big.Int).Add(result, key)
		t.Add(t, big.NewInt(int64(c)))
		t.Mod(t, p)
		// result = t^3 mod p
		t2 := new(big.Int).Mul(t, t)
		t2.Mod(t2, p)
		result = new(big.Int).Mul(t2, t)
		result.Mod(result, p)
	}
	// ciphertext = result + key mod p
	result.Add(result, key)
	result.Mod(result, p)
	return result
}
