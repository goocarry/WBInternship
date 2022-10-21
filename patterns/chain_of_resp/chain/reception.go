package chain

import "fmt"

// Reception ...
type Reception struct {
	Next Department
}

// Execute ...
func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient already registered")
		r.Next.Execute(p)
		return
	}

	fmt.Println("Registering pacient")
	p.RegistrationDone = true
	r.Next.Execute(p)
}

// SetNext ...
func (r *Reception) SetNext(dep Department) {
	r.Next = dep
}