package main

import "fmt"

func main() {
	arr := []int{2, 10, 9, 15, 3}

	sort := sort(arr)
	fmt.Println(sort)
}

func sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}

	for i, v := range arr {
		if pivot < arr[i] {
			right = append(right, v)
		}
		if pivot > arr[i] {
			left = append(left, v)
		}
	}
	left = sort(append(left, pivot))
	right = sort(right)
	result := append(left, right...)
	return result

}
