package main

import "fmt"

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sortMap := make(map[int][]float64)

	// -25.4 => -20
	// -13.5 => -10
	// -9.13  => 0
	// 2.25 => 10
	// 13.44 => 20
	for i, v := range temps {
		fmt.Printf("i=%d, v=%f\n", i, v)
		if v > 0 && v < 10 {
			arr, ok := sortMap[i]
			if ok {
				arr = append(arr, v)
			} else {
				arr = make([]float64, 0)
				arr = append(arr, v)
			}
			sortMap[0] = arr
		} else {
			a := int(v/10) * 10
			arr, ok := sortMap[a]
			if ok {
				arr = append(arr, v)
			} else {
				arr = make([]float64, 0)
				arr = append(arr, v)
			}
			sortMap[a] = arr
		}
	}
	fmt.Printf("sorted: %v", sortMap)
}
