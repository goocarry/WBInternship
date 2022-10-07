package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter ...
type Counter struct {
	value int
}

func (c *Counter) increment() {
	c.value++
}

func main() {
	counter := Counter{value: 0}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("1 goroutine")
			counter.increment()
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("2 goroutine")
			counter.increment()
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(counter.value)
}
