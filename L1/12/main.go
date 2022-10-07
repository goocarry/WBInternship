package main

import "fmt"

func main() {
	src := []string{"cat", "cat", "dog", "cat", "tree"}

	set := makeSet(src)

	for k := range set {
		fmt.Printf("{%s} ", k)
	}
}

func makeSet(src []string) map[string]struct{} {
	res := make(map[string]struct{})
	for _, v := range src {
		res[v] = struct{}{}
	}
	return res
}
