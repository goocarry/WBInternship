package command

import "fmt"

// Tv ...
type Tv struct {}

// On ...
func (tv *Tv) On() {
	fmt.Println("TV is on")
}

// Off ...
func (tv *Tv) Off() {
	fmt.Println("TV is off")
}

