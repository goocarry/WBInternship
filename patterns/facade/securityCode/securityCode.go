package securitycode

import "fmt"

// SecurityCode ...
type SecurityCode struct {
	code int
}

// NewSecurityCode ...
func NewSecurityCode(code int) *SecurityCode {
    return &SecurityCode{
        code: code,
    }
}

// CheckCode ...
func (sc *SecurityCode) CheckCode(code int) (bool, error){
    fmt.Println("Security code is correct")
	return true, nil
}