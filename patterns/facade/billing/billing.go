package billing

import "fmt"

// Billing ...
type Billing struct {
	transactions []string
}

// NewBilling ...
func NewBilling() *Billing {
    return &Billing{
        transactions: []string{},
    }
}

// SaveTransaction ...
func (b *Billing) SaveTransaction(trID string) {
	b.transactions = append(b.transactions, trID)
    fmt.Println("Transaction saved successfully")
}

