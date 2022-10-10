package main

import "fmt"

func main() {
	string1 := "abcd"
	string2 := "abCdefAaf"
	string3 := "aabcd"

	fmt.Println(checkUnique(string1))
	fmt.Println(checkUnique(string2))
	fmt.Println(checkUnique(string3))
}

func checkUnique(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		if _, ok := seen[r]; ok {
			return false
		} 
		seen[r] = true
	}
	return true
}