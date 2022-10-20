package builder

// SedanBuilder ...
type SedanBuilder struct {
	BodyType string
	DriveType string
	EngineType float64
}

// NewSedanBuilder ...
func NewSedanBuilder() *SedanBuilder {
	return &SedanBuilder{}
}

// SetBody ...
func (sb *SedanBuilder) SetBody() {
	sb.BodyType = "sedan"
}

// SetDrive ...
func (sb *SedanBuilder) SetDrive() {
	sb.DriveType = "front wheel drive"
}

// SetEngine ...
func (sb *SedanBuilder) SetEngine() {
	sb.EngineType = 1.8
}

// GetCar ...
func (sb *SedanBuilder) GetCar() Car {
	return Car{
		BodyType: sb.BodyType,
		DriveType: sb.DriveType,
		EngineType: sb.EngineType,
	}
}