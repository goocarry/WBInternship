package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := point{x: 1, y: 1}
	p2 := point{x: 4, y: 10}
	distance:= distance(p1, p2)
	fmt.Println("distance: ", distance)
}

type point struct {
	x int
	y int
}

func newPoint(x int, y int) *point {
	return &point{x: x, y: y}
}

func distance(p1 point, p2 point) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y
	return math.Sqrt(float64(x*x + y*y))	
}
 
