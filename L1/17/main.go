package main

import "fmt"

func main() {
	// only for sorted array
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	num := 8

	index := binarySearch(nums, num, 0, len(nums)-1)
	
	if index > 0 {
		fmt.Println("index", index)
	} else {
		fmt.Println("number not found")
	}
}

// iterative method
func binarySearch(nums []int, num int, low int, high int) int {

	for low <= high {
		middle := (low + high) / 2
		if nums[middle] == num {
			return middle
		} else if nums[middle] < num {
			low = middle + 1
			binarySearch(nums, num, low, high)
		} else {
			high = middle - 1
			binarySearch(nums, num, low, high)
		}
	}
	return -1
}
