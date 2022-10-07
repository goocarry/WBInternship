package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	split := strings.Split(str, " ")
	fmt.Println(split)

	result := ""

	for _, v := range split {
		result = string(v) + " " + result
	}

	fmt.Println(result)
}
