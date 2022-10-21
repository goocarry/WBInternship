package visitor

// Triangle ...
type Triangle struct {
	A int
	B int
	C int
}

// Accept ...
func (t *Triangle) Accept(visitor ShapeVisitor) {
	visitor.VisitTriangle(t)
}

// GetType ...
func (t *Triangle) GetType() string {
	return "Triangle"
}