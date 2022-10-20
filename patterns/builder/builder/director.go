package builder

// Director ...
type Director struct {
	builder CarBuilder
}

// NewDirector ...
func NewDirector(builder CarBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

// SetBuilder ...
func (d *Director) SetBuilder(builder CarBuilder) {
	d.builder = builder
}

// BuildCar ...
func (d *Director) BuildCar() Car {
	d.builder.SetBody()
	d.builder.SetDrive()
	d.builder.SetEngine()
	return d.builder.GetCar()
}