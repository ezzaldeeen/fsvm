package engine

import (
	"fmt"
)

// Depositing ...
type Depositing struct {
	engine *Engine
}

// Deposit represents the depositing action on this state
func (s Depositing) Deposit(amount float64) error {
	err := s.engine.addToBalance(amount)
	if err != nil {
		return err
	}
	// check if the entered amount is enough to dispense
	total := s.engine.getTotalPrice()
	balance := s.engine.balance
	if balance >= total {
		s.engine.setState(s.engine.dispensing)
		return nil
	}
	// compute the remaining amount of money for dispensing
	remaining := total - balance
	return fmt.Errorf(
		"[DEPOSITING] - reamining: %0.2f$ to dispense", remaining,
	)
}

// Cancel represents the canceling action on this state
func (s Depositing) Cancel() error {
	// todo: dispense money
	s.engine.reset()
	s.engine.setState(s.engine.selecting)
	return nil
}

// Select represents the selecting action on this state
func (s Depositing) Select(_ string) error {
	return fmt.Errorf( // todo: what if we want to select multiple items?
		"[DEPOSITING] - already selected, please deposit to dispense",
	)
}

// Dispense represents the dispensing action on this state
func (s Depositing) Dispense() error {
	return fmt.Errorf(
		"[DEPOSITING] - unable to dispense, please deposit first",
	)
}
