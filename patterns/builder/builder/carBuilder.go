package builder

// CarBuilder ...
type CarBuilder interface {
	SetBody()
	SetDrive()
	SetEngine()
	GetCar() Car
}

// GetCarBuilder ...
func GetCarBuilder(builderType string) CarBuilder {
	if builderType == "sedan" {
		return NewSedanBuilder()
	}

	if builderType == "suv" {
		return NewSuvBuilder()
	}

	return nil
}