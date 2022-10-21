package visitor

// Square ...
type Square struct {
	Side int
}

// Accept ...
func (s *Square) Accept(visitor ShapeVisitor) {
	visitor.VisitSquare(s)
}

// GetType ...
func (s *Square) GetType() string {
	return "Square"
}