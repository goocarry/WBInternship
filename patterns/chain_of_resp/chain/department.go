package chain

// Department ...
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}