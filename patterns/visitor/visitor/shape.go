package visitor

// Shape ...
type Shape interface {
	Accept(ShapeVisitor) 
	GetType() string
}