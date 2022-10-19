package main

import (
	"fmt"

	"github.com/goocarry/wb-internship/patterns/fabric/gun"
)

func main() {
	fmt.Println("fabric example")

	ak47, _ := gun.GetGun("ak47")
    musket, _ := gun.GetGun("musket")

	fmt.Println(ak47.GetName())
	fmt.Println(musket.GetName())

}