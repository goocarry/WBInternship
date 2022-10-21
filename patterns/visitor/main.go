package main

import (
	"fmt"

	"github.com/goocarry/wb-internship/patterns/visitor/visitor"
)

func main() {
	fmt.Println("visitor example")

	circle := &visitor.Circle{
		Radius: 5,
	}
	square := &visitor.Square{
		Side: 4,
	}

	areaCalc := &visitor.AreaCalculator{}

	circle.Accept(areaCalc)
	square.Accept(areaCalc)
}