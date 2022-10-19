package gun

// Gun ...
type Gun struct {
	name string
	power int
}


// SetName ...
func (g *Gun) SetName(name string) {
	g.name = name
}

// SetPower ...
func (g *Gun) SetPower(power int) {
	g.power = power
}

// GetName ...
func (g *Gun) GetName() string {
	return g.name 
}

// GetPower ...
func (g *Gun) GetPower()  int {
	return g.power
}