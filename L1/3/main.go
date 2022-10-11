package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var result int
	var wg sync.WaitGroup

	a := []int{2, 4, 6, 8, 10}
	resSlice := make([]int, 5)

	wg.Add(len(a))

	for i := range a {
		// use goroutine for concurrent func call
		go square(a[i], resSlice, i, &wg)
	}
	wg.Wait()

	fmt.Fprintf(os.Stdout, "result: %v\n", resSlice)

	for j := range resSlice {
		result += resSlice[j]
	}
	fmt.Fprintf(os.Stdout, "result: %d\n", result)

}

func square(i int, resultSlice []int, pos int, wg *sync.WaitGroup) {
	resultSlice[pos] = i * i
	wg.Done()
}
