package machine

import "errors"

type Cancel struct {
	machine *Machine
}

func (c Cancel) Select(_ string) error {
	return errors.New(
		"unable to select, canceling in progress",
	)
}

func (c Cancel) Deposit(_ float64) error {
	return errors.New(
		"unable to select, canceling in progress",
	)
}

func (c Cancel) Dispense() error {
	//TODO implement me
	panic("implement me")
}

func (c Cancel) Cancel() error {
	//TODO implement me
	panic("implement me")
}
