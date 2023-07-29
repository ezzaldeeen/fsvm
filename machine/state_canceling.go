package machine

import "errors"

// Canceling ...
type Canceling struct {
	machine *Machine
}

// Select represents the selecting action on this state
func (c Canceling) Select(_ string) error {
	return errors.New(
		"unable to select, canceling in progress",
	)
}

// Deposit represents the depositing action on this state
func (c Canceling) Deposit(_ float64) error {
	return errors.New(
		"unable to select, canceling in progress",
	)
}

// Dispense represents the dispensing action on this state
func (c Canceling) Dispense() error {
	//TODO implement me
	panic("implement me")
}

// Cancel represents the canceling action on this state
func (c Canceling) Cancel() error {
	//TODO implement me
	panic("implement me")
}
