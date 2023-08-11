package engine

import (
	"fmt"
)

// Selecting ...
type Selecting struct {
	engine *Engine
}

// Select represents the selecting action on this state
func (s Selecting) Select(itemID string) error {
	err := s.engine.addToBasket(itemID)
	if err != nil {
		// todo: rewrite this
		return fmt.Errorf("[SELECTING] - %s", err)
	}
	s.engine.setState(s.engine.depositing)
	return nil
}

// Deposit represents the depositing action on this state
func (s Selecting) Deposit(_ float64) error {
	return fmt.Errorf( // TODO: can deposit even not select an item
		"[SELECTING] - you didn't select to deposit, please select an item first",
	)
}

// Dispense represents the dispensing action on this state
func (s Selecting) Dispense() error {
	return fmt.Errorf(
		"[SELECTING] - you didn't select to dispense, please select an item first",
	)
}

// Cancel represents the canceling action on this state
func (s Selecting) Cancel() error {
	return fmt.Errorf(
		"[SELECTING] - you didn't select to cancel, please select an item first",
	)
}
