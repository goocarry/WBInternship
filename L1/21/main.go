package main

import "fmt"

// Animal ...
type Animal interface {
	sayGav()
}

// Dog ...
type Dog struct{}

func (d *Dog) sayGav() {
	fmt.Println("gav")
}

// Cat ...
type Cat struct{}

func (c *Cat) sayMyau() {
	fmt.Println("myau")
}

// CatAdapter ...
type CatAdapter struct {
	cat *Cat
}

func (ca *CatAdapter) sayGav() {
	ca.cat.sayMyau()
}

// Human ...
type Human struct{}

func (h *Human) makeSay(a Animal) {
	a.sayGav()
}

func main() {
	human := &Human{}
	dog := &Dog{}
	human.makeSay(dog)
	cat := &Cat{}
	adapter := &CatAdapter{
		cat: cat,
	}
	human.makeSay(adapter)
}
