package chain

import "fmt"

// Doctor ...
type Doctor struct {
	Next Department
}

// Execute ...
func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.Next.Execute(p)
		return	
	}

	fmt.Println("Doctor checkup")
	p.DoctorCheckUpDone = true
	d.Next.Execute(p)
}

// SetNext ...
func (d *Doctor) SetNext(dep Department) {
	d.Next = dep
}