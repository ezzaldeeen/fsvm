package machine

import (
	"fmt"
	"log"
	"math"
	"time"
)

type Depositing struct {
	machine *Machine
}

func (s Depositing) Select(_ string) error {
	return fmt.Errorf("item already selected, please deposit")
}

func (s Depositing) Deposit(amount float64) error {
	err := s.machine.AddToBalance(amount)
	if err != nil {
		return err
	}
	price := s.machine.GetTotalPrice()
	balance := s.machine.GetBalance()
	remaining := math.Abs(price - balance)
	if balance < price {
		return fmt.Errorf("money not enough, remaining: %.02f$", remaining)
	}
	log.Printf("money diposeted successfuly, ready for dispinsing\n")
	err = s.machine.SetBalance(remaining)
	if err != nil {
		return err
	}
	log.Printf("balance: %.02f$\n", remaining)
	s.machine.SetState(s.machine.GetDispensingState())
	return nil
}

func (s Depositing) Dispense() error {
	return fmt.Errorf("unable to dispense, deposit first")

}

func (s Depositing) Cancel() error {
	log.Println("Canceling...")
	time.Sleep(1 * time.Second)
	log.Printf(
		"Money: %.02f$ dispensed\n",
		s.machine.GetBalance(),
	)
	s.machine.CleanBasket()
	err := s.machine.SetBalance(0)
	if err != nil {
		return err
	}
	s.machine.SetState(s.machine.GetSelectingState())
	return nil
}