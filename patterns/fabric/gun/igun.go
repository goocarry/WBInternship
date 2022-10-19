package gun

// IGun ...
type IGun interface {
	SetName(string)
	SetPower(int)
	GetName() string
	GetPower() int
}