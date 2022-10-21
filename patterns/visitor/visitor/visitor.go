package visitor

// ShapeVisitor ...
type ShapeVisitor interface {
	VisitSquare(*Square)
	VisitTriangle(*Triangle)
	VisitCircle(*Circle)
}