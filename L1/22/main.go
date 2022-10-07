package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("123000000000000000000000000000000", 10)
	b, _ := new(big.Int).SetString("234000000000000444400000000000000", 10)

	fmt.Printf("add: %v\n", Add(a, b))
	fmt.Printf("sub: %v\n", Sub(a, b))
	fmt.Printf("mul: %v\n", Mul(a, b))
	fmt.Printf("div: %v\n", Div(a, b))
}

// Add ...
func Add(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// Sub ...
func Sub(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// Mul ...
func Mul(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// Div ...
func Div(a *big.Int, b *big.Int) *big.Float {
	f1 := new(big.Float).SetInt(a)
	f2 := new(big.Float).SetInt(b)
	return new(big.Float).Quo(f1, f2)
}
