package main

import "fmt"

func main() {
	// simply convering bytes to string and concat
	// this solution can be a bottleneck
	// we can use slice of runes
	// or strings.Builder
	str := "абырвалг"
	result := ""

	for _, v := range str {
		result = string(v) + result
	}

	fmt.Println(result)
}
