package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeInSec := 5
	ch := make(chan int)
	defer close(ch)

	go reader(ch)

	for timeInSec != 0 {
		ch <- rand.Intn(100)
		timeInSec--
		time.Sleep(1 * time.Second)
	}
}

func reader(ch <-chan int) {
	for msg := range ch {
		fmt.Printf("new msg: %d\n", msg)
	}
}
