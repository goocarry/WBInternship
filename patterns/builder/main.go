package main

import (
	"fmt"

	"github.com/goocarry/wb-internship/patterns/builder/builder"
)

func main() {
	fmt.Println("builder example")
	
	suvBuilder := builder.GetCarBuilder("suv")
	sedanBuilder := builder.GetCarBuilder("sedan")

	director := builder.NewDirector(suvBuilder)
	suv := director.BuildCar()

	fmt.Printf("suv body: %s\n", suv.BodyType)

	director.SetBuilder(sedanBuilder)
	sedan := director.BuildCar()

	fmt.Printf("sedan body: %s\n", sedan.BodyType)
}