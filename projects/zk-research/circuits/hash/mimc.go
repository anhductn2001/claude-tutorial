// Package hash implements ZK circuit gadgets for hash functions.
package hash

import (
	"github.com/consensys/gnark/frontend"
)

// mimcRounds is the number of rounds for this demo circuit.
const mimcRounds = 7

// mimcConstants are simple sequential round constants for this demo.
// Production MiMC uses field-specific constants generated from a hash-to-field step.
var mimcConstants = [mimcRounds]int{0, 1, 2, 3, 4, 5, 6}

// MiMCCircuit proves knowledge of a preimage x such that MiMC(x, k) == hash.
//
// Public inputs:  Hash, Key
// Private witness: Preimage
type MiMCCircuit struct {
	// Public inputs
	Hash frontend.Variable `gnark:",public"`
	Key  frontend.Variable `gnark:",public"`

	// Private witness
	Preimage frontend.Variable
}

// Define declares the constraints for the MiMC circuit.
//
// MiMC encryption: E_k(x) = (x + k + c_i)^3 iterated r rounds, then add k.
// We use 7 rounds here for demonstration. Production uses 110 rounds.
func (c *MiMCCircuit) Define(api frontend.API) error {
	result := c.Preimage

	for _, rc := range mimcConstants {
		// t = result + key + round_constant
		t := api.Add(result, c.Key, rc)

		// t^3 = t * t * t
		t2 := api.Mul(t, t)
		result = api.Mul(t2, t)
	}

	// Final key addition: ciphertext = result + key
	ciphertext := api.Add(result, c.Key)

	// Constrain: ciphertext must equal the claimed hash
	api.AssertIsEqual(ciphertext, c.Hash)

	return nil
}
