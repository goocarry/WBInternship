package main

import "fmt"

func main() {
	nums := []int{2, 4, 5, 6, 7, 8}
	i := 2

	result := append(nums[:i], nums[i+1:]...)

	fmt.Println(result)
}
