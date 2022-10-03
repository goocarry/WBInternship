package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	initialCh := make(chan int)
	squareCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		i := 0
		for {
			initialCh <- i
			i++
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for num := range initialCh {
			squareCh <- num * num
		}
	}()

	go func() {
		for squreNum := range squareCh {

			fmt.Printf("square: %d\n", squreNum)
		}
	}()

	wg.Wait()

	defer close(initialCh)
	defer close(squareCh)
}
