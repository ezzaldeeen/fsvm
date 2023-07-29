package machine

import (
	"fmt"
	"log"
)

type Selecting struct {
	machine *Machine
}

// Select represents the selecting action on this state
func (s Selecting) Select(itemID string) error {
	if err := s.machine.AddToBasket(itemID); err != nil {
		return err
	}
	log.Printf("item [%s] added to basket", itemID)
	s.machine.SetState(s.machine.GetDepositingState())
	return nil
}

// Deposit represents the depositing action on this state
func (s Selecting) Deposit(_ float64) error {
	return fmt.Errorf(
		"unable to deposit, please select item first\n",
	)
}

// Dispense represents the dispensing action on this state
func (s Selecting) Dispense() error {
	return fmt.Errorf(
		"unable to dispense, please select item first\n",
	)
}

// Cancel represents the canceling action on this state
func (s Selecting) Cancel() error {
	return fmt.Errorf(
		"unable to cancel, please select item first\n",
	)
}
