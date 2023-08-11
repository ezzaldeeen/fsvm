package engine

import (
	"fmt"
)

// Dispensing ...
type Dispensing struct {
	engine *Engine
}

// Dispense represents the dispensing action on this state
func (s Dispensing) Dispense() error {
	// calculate change to dispense
	total := s.engine.getTotalPrice()
	balance := s.engine.balance
	change := balance - total
	fmt.Printf("[DISPENSING] - money dispensed: %0.2f\n", change)
	//
	for _, item := range s.engine.basket {
		item.quantity--
		fmt.Printf("[DISPENSING] - item dispensed: %s\n", item.name)
	}

	s.engine.reset()
	s.engine.setState(s.engine.selecting)
	return nil
}

// Select represents the selecting action on this state
func (s Dispensing) Select(_ string) error {
	return fmt.Errorf(
		"[DISPENSING] - unable to select, dispensing in progress",
	)
}

// Deposit represents the depositing action on this state
func (s Dispensing) Deposit(_ float64) error {
	return fmt.Errorf(
		"[DISPENSING] - unable to deposit, dispensing in progress",
	)
}

// Cancel represents the canceling action on this state
func (s Dispensing) Cancel() error {
	return fmt.Errorf(
		"[DISPENSING] - unable to cancel, dispensing in progress",
	)
}
