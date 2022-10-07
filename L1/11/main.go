package main

import "fmt"

// Empty struct
var Empty struct{}

func main() {
	numsA := []int{2, 4, 5, 7, 3, 6, 7}
	numsB := []int{2, 5, 11, 15, 6, 12}

	fmt.Println("nums a :", numsA)
	fmt.Println("nums b :", numsB)

	result := make(map[int]bool)

	for _, v := range numsA {
		result[v] = false
	}

	for _, v := range numsB {
		if _, ok := result[v]; ok {
			result[v] = true
		}
	}
	for k, v := range result {
		if v != true {
			delete(result, k)
		}
	}
	fmt.Printf("%v", result)

}
