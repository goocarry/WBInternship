package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var workersNum *int

func init() {
	// get number of workers from flags on start
	workersNum = flag.Int("w", 3, "number of workers")
	flag.Parse()
}

func worker(i int, ch chan int) {
	// range operator prints messages from channel until channel is closed
	// we use 'defer close(ch)' and after that all the workers will exit
	for msg := range ch {
		fmt.Printf("worker #%d: %d\n", i, msg)
	}
}

func main() {

	fmt.Printf("workers num: %d\n", *workersNum)

	ch := make(chan int)
	defer close(ch)

	for i := 0; i < *workersNum; i++ {
		go worker(i, ch)
	}

	// write to channel every 2 sec until ctrl+c
	for true {
		ch <- rand.Intn(10)
		time.Sleep(2 * time.Second)
	}
}
