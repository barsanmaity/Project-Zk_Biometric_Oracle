package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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
    //json bridge
	fmt.Println("reading json file..")
	file, err:= os.Open("../processor_engine/fingerprint.json")
	if err != nil {
		log.Fatal("not file open: ", err)
	}
	defer file.Close()

    var ectractedNumbers[13] int
	if err := json.NewDecoder(file).Decode(&ectractedNumbers);
	err != nil {
		log.Fatal("could not parse json: ", err)
	}

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

    //save keys to disk
	pkFile, _ :=os.Create("provingKey.bin")
	provingKey.WriteTo(pkFile)
	pkFile.Close()
    
	vkFile, _ :=os.Create("verifyingKey.bin")
	verifyingKey.WriteTo(vkFile)
	vkFile.Close()

	fmt.Println("injecting fingerprint..")
	var witnessVars [13]frontend.Variable
	for i := 0; i<13; i++ {
          witnessVars[i] = ectractedNumbers[i]
	}

	assignment := Vcircuit{
		Fingerprint: witnessVars,
	}
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		log.Fatal("witness create error: ", err)
	}
	//generating proof 
	fmt.Println("generating proof..")
	proof, err := groth16.Prove(compileConstraint, provingKey, witness)
	if err != nil {
		log.Fatal("Prove error: ", err)
	}
	fmt.Println("saving proof in disk")
	prooffile, _ :=os.Create("proof.bin")
	proof.WriteTo(prooffile)
	prooffile.Close()
	
	//verify
	publicWitness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField(), frontend.PublicOnly())
	if err != nil {
		log.Fatal("public witness error: ", err)
	}

	err = groth16.Verify(proof, verifyingKey, publicWitness)
	if err != nil {
		log.Fatal("verification failed: ", err)
	}
	fmt.Println("verification success, proof complete ")
}