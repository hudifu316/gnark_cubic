package main

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func main() {
	fmt.Println("Cubic sample")
	var myCircuit Circuit
	r1cs, _ := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &myCircuit)

	//1.Setup
	fmt.Println("1.Setup")
	pk, vk, _ := groth16.Setup(r1cs)

	//2.Prove
	fmt.Println("2.Prove")
	assignment := &Circuit{
		X: 3,
		Y: 35,
	}
	witness, err := frontend.NewWitness(assignment, ecc.BN254)
	if err != nil {
		fmt.Println("Prove Err: ", err)
	}
	fmt.Println("Witness: ", witness)

	proof, _ := groth16.Prove(r1cs, pk, witness)
	fmt.Println("proof: ", proof)

	//3.Verify
	fmt.Println("3.Verify")
	publicWitness, _ := witness.Public()
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		fmt.Println("Verify Err: ", err)
	}
}
