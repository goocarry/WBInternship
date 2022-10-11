package main

import "fmt"

func main() {
	var i, j = 1, 2
	i, j = swap(i, j)
	fmt.Printf("i: %d, j: %d\n", i, j)

	var a, b = 3, 4
	b, a = swap2(a, b)
	fmt.Printf("a: %d, b: %d\n", a, b)
}

func swap(i int, j int) (int, int) {
	return j, i
}

func swap2(i int, j int) (int, int) {
	j = i + j
	i = j - i
	j = j - i
	return j, i
}
