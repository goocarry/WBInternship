package main

import "fmt"

// Human ...
type Human struct {
	Name string
}

// Action is struct with Human's struct methods and fields
type Action struct {
	Human
}

func (h *Human) sayName() {
	fmt.Printf("my name is %s", h.Name)
}

func main() {
	a := Action{}
	a.Name = "Bob"
	a.sayName()
}
