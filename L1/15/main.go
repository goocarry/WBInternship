package main

import (
	"fmt"
	"math/rand"
)

var justString string

func main() {
	someFunc()
}

func someFunc() {
	v := createHugeString(1 << 10)

	// we have to check the length of the slice
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		fmt.Println("out of slice bounds")
		justString = ""
	}

	fmt.Println(justString)
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// make a guess that createHugeString recieves size of string
func createHugeString(size int) string {
	randomString := make([]byte, size)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomString)
}
