package visitor

import "fmt"

// AreaCalculator ...
type AreaCalculator struct {
	area int
}

// VisitCircle ...
func (a *AreaCalculator) VisitCircle(c *Circle) {
	fmt.Println("Calculate area for circle")
}


// VisitSquare ...
func (a *AreaCalculator) VisitSquare(s *Square) {
	fmt.Println("Calculate area for square")
}


// VisitTriangle ...
func (a *AreaCalculator) VisitTriangle(t *Triangle) {
	fmt.Println("Calculate area for triangle")
}