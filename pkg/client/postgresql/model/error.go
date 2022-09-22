package model

import "fmt"

// ErrCommit ...
func ErrCommit(err error) error {
	return fmt.Errorf("failed to commit Tx due to error: #{err}")
}

// ErrRollback ...
func ErrRollback(err error) error {
	return fmt.Errorf("failed to rollback Tx due to error: #{err}")
}

// ErrCreateTx ...
func ErrCreateTx(err error) error {
	return fmt.Errorf("failed to create Tx due to error: #{err}")
}

// ErrCreateQuery ...
func ErrCreateQuery(err error) error {
	return fmt.Errorf("failed to create SQL Query due to error: #{err}")
}

// ErrScan ...
func ErrScan(err error) error {
	return fmt.Errorf("failed to scan due to error: #{err}")
}

// ErrDoQuery ...
func ErrDoQuery(err error) error {
	return fmt.Errorf("failed to query due to error: #{err}")
}
