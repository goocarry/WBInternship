package chain

import "fmt"

// Cashier ...
type Cashier struct {
	Next Department
}

// Execute ...
func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment already done")
	}
	fmt.Println("Payment")
}

// SetNext ...
func (c *Cashier) SetNext(dep Department) {
	c.Next = dep
}