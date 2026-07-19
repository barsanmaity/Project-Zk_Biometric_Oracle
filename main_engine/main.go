package main

import (
	"fmt"
	"log"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)


//zk-circuit for the MFCC
type Vcircuit struct {
	//hold fingerprint
	Fingerprint [13]frontend.Variable `gnark:",secret"` //hidden in final proof
}

//math rules
func (circuit *Vcircuit) Define(api frontend.API) error {
	//none of the fingerprint value != 0
	for i:= 0; i<13; i++ {
		api.AssertIsDifferent(circuit.Fingerprint[i], 0)
	}
	return nil
}

func main() {
	fmt.Println("ZK Engine Initialized!")
	var circuit Vcircuit

	//compile circuit to R1CS
    fmt.Println("compile the voice fingerprint")
	compileConstraint, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	 if err != nil {
		log.Fatal("compilation error: ", err)
	 }
	 fmt.Println("circuit compiled successfully")
	 fmt.Println("Total math constraints: ", compileConstraint.GetNbConstraints())

	 //run setup and proving and generate keys
	 fmt.Println("generating keys...")
	 provingKey, verifyingKey , err := groth16.Setup(compileConstraint)
	 if err != nil {
		log.Fatal("setup error: ", err)
	 }
	 fmt.Println("Proving key and Verifying key generated successfully")

	//silence unused variables
	_ = provingKey
	_ = verifyingKey	
}