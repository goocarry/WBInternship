package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println(time.Now())
	sleep1(2)
	fmt.Println(time.Now())
	sleep2(4)
	fmt.Println(time.Now())
	sleep3(2)
	fmt.Println(time.Now())
}

// with time.Now()
func sleep1(seconds int) {
	start := time.Now().Unix()
	for time.Now().Unix()-start < int64(seconds) {}
}

// with time.After
func sleep2(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}

// with timer
func sleep3(seconds int) {
	timer := time.NewTimer(time.Second * time.Duration(seconds) )
	<-timer.C
}