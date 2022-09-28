package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	for i := range a {
		// use goroutine for concurrent func call
		go square(a[i])
	}
	time.Sleep(5 * time.Second)
}

func square(i int) int {
	fmt.Fprintf(os.Stdout, "%d\n", i*i)
	return i * i
}
