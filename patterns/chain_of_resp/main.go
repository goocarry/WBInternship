package main

import (
	"fmt"

	"github.com/goocarry/wb-internship/patterns/chain_of_resp/chain"
)

func main() {
	fmt.Println("chain of responsibility example")

	cachier := &chain.Cashier{}

	doctor := &chain.Doctor{}
	doctor.SetNext(cachier)

	rec := &chain.Reception{}
	rec.SetNext(doctor)

	patient := &chain.Patient{Name: "Ivan"}
	rec.Execute(patient)

}