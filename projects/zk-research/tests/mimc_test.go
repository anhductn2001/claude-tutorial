package tests

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/test"

	"github.com/decentrio/zk-research/circuits/hash"
)

// TestMiMC_ValidWitness checks that a correct preimage produces an accepted proof.
func TestMiMC_ValidWitness(t *testing.T) {
	circuit := &hash.MiMCCircuit{}

	preimage := big.NewInt(42)
	key := big.NewInt(0)
	expectedHash := computeMiMCNative(preimage, key)

	assignment := &hash.MiMCCircuit{
		Preimage: preimage,
		Key:      key,
		Hash:     expectedHash,
	}

	assert := test.NewAssert(t)
	assert.ProverSucceeded(circuit, assignment, test.WithCurves(ecc.BN254))
}

// TestMiMC_InvalidWitness checks that a wrong preimage is rejected.
func TestMiMC_InvalidWitness(t *testing.T) {
	circuit := &hash.MiMCCircuit{}

	// Use an intentionally wrong hash value
	assignment := &hash.MiMCCircuit{
		Preimage: frontend.Variable(42),
		Key:      frontend.Variable(0),
		Hash:     frontend.Variable(9999), // wrong hash
	}

	assert := test.NewAssert(t)
	assert.ProverFailed(circuit, assignment, test.WithCurves(ecc.BN254))
}

// TestMiMC_Groth16_EndToEnd runs the full compile → setup → prove → verify pipeline.
func TestMiMC_Groth16_EndToEnd(t *testing.T) {
	var circuit hash.MiMCCircuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		t.Fatalf("compile: %v", err)
	}

	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	expectedHash := computeMiMCNative(big.NewInt(42), big.NewInt(0))
	assignment := hash.MiMCCircuit{
		Preimage: big.NewInt(42),
		Key:      big.NewInt(0),
		Hash:     expectedHash,
	}

	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		t.Fatalf("witness: %v", err)
	}

	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		t.Fatalf("prove: %v", err)
	}

	publicWitness, err := witness.Public()
	if err != nil {
		t.Fatalf("public witness: %v", err)
	}

	if err := groth16.Verify(proof, vk, publicWitness); err != nil {
		t.Fatalf("verify: %v", err)
	}
}

// bn254ScalarField is the BN254 scalar field modulus.
var bn254ScalarField, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088548364400416034343698204186575808495617", 10,
)

// computeMiMCNative mirrors the circuit's 7-round MiMC using big.Int field arithmetic.
func computeMiMCNative(preimage, key *big.Int) *big.Int {
	p := bn254ScalarField
	result := new(big.Int).Set(preimage)
	for c := 0; c < 7; c++ {
		t := new(big.Int).Add(result, key)
		t.Add(t, big.NewInt(int64(c)))
		t.Mod(t, p)
		t2 := new(big.Int).Mul(t, t)
		t2.Mod(t2, p)
		result = new(big.Int).Mul(t2, t)
		result.Mod(result, p)
	}
	result.Add(result, key)
	result.Mod(result, p)
	return result
}
