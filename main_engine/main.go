package main

import (
	"fmt"
	"github.com/consensys/gnark/frontend"
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
	fmt.Println("Blueprint for the 13-element Voice Circuit is ready.")
}