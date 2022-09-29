package main

import (
	"fmt"
	"sync"
)

func main() {

	// possible ways to quit goroutine
	// 1 way is to pass signal channel
	quit := make(chan bool)
	fmt.Println("start first goroutine")
	go func(quit <-chan bool) {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("goroutine still works")
			}

		}
	}(quit)

	// quit 1st goroutine
	quit <- true
	fmt.Println("first goroutine done")

	// 2 way is to use waitGroup
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("start second goroutine")
	go func() {
		fmt.Println("print and call wg.Done()")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("second goroutine done")

	// 3 way is to close channel passed to goroutine
	ch := make(chan bool)
	fmt.Println("start third goroutine")
	go func(ch <-chan bool) {
		for msg := range ch {
			fmt.Printf("goroutine works until channel is closed: %t", msg)
		}
	}(ch)

	// quit 3rd goroutine
	close(ch)
	fmt.Println("third goroutine done")
}
