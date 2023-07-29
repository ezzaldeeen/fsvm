package machine

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

type Dispensing struct {
	machine *Machine
}

func (s Dispensing) Select(_ string) error {
	return fmt.Errorf(
		"unable to select, dispensing in progress",
	)
}

func (s Dispensing) Deposit(_ float64) error {
	return fmt.Errorf(
		"unable to deposit, dispensing in progress",
	)
}

func (s Dispensing) Dispense() error {
	log.Println("Dispensing...")
	time.Sleep(1 * time.Second)
	log.Printf(
		"item: %v dispensed - Money: %.02f$ dispensed\n",
		s.machine.basket,
		s.machine.balance,
	)
	err := s.machine.SetBalance(0)
	if err != nil {
		return err
	}
	// todo: compose the below functions into single function
	s.machine.CleanBasket()
	s.machine.sessionID = uuid.NewString()
	s.machine.SetState(s.machine.GetSelectingState())
	return nil
}

func (s Dispensing) Cancel() error {
	return fmt.Errorf(
		"unable to cancel, dispensing in progress",
	)
}
