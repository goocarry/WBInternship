package builder

// SuvBuilder ...
type SuvBuilder struct {
	BodyType string
	DriveType string
	EngineType float64
}

// NewSuvBuilder ...
func NewSuvBuilder() *SuvBuilder {
	return &SuvBuilder{}
}

// SetBody ...
func (sb *SuvBuilder) SetBody() {
	sb.BodyType = "SUV"
}

// SetDrive ...
func (sb *SuvBuilder) SetDrive() {
	sb.DriveType = "4x4 drive"
}

// SetEngine ...
func (sb *SuvBuilder) SetEngine() {
	sb.EngineType = 3.0
}

// GetCar ...
func (sb *SuvBuilder) GetCar() Car {
	return Car{
		BodyType: sb.BodyType,
		DriveType: sb.DriveType,
		EngineType: sb.EngineType,
	}
}