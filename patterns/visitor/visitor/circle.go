package visitor

// Circle ...
type Circle struct {
	Radius int
}

// Accept ...
func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

// GetType ...
func (c *Circle) GetType() string {
	return "Circle"
}