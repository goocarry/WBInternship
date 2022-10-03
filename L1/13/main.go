package main

import "fmt"

// we can use swap function
// OR
// we can use XOR operator
// OR
// we can use pointers
func main() {
	var i, j = 1, 2
	i, j = swap(i, j)
	fmt.Printf("i: %d, j: %d", i, j)
}

func swap(i int, j int) (int, int) {
	return j, i
}
