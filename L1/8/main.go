package main

import "fmt"

func main() {
	test := setBit(256, 5)
	fmt.Printf("new number %d", test)
}

func setBit(n int64, i int) int64 {
	var mask int64
	mask = 1 << i // if n = 5 then mask will be 000....10000

	result := n | mask // logical OR will apply mask witn n'th bit
	return result
}
