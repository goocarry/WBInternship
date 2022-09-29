package main

import (
	"fmt"
	"sync"
)

var m map[string]string

func main() {
	timeInSec := 10
	mutex := &sync.Mutex{}
	m = make(map[string]string)
	var wg sync.WaitGroup

	for i := 0; i < timeInSec; i++ {
		wg.Add(1)
		go writer(i, mutex, &wg)
	}

	wg.Wait()

	fmt.Printf("result map %v:", m)
}

func writer(i int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	m[fmt.Sprint(i)] = "writer #" + fmt.Sprint(i)
	fmt.Printf("writer #%d\n", i)
	mu.Unlock()
	wg.Done()
}
