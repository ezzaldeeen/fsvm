package machine

import (
	"fmt"
	"github.com/ezzaldeeen/fsvm/eventstore"
	"github.com/google/uuid"
)

// Item represents an item in the vending machine.
type Item struct {
	price    float64
	quantity uint
}

// Machine represents a vending machine.
// It maintains the current state, balance, basket, and item inventory.
type Machine struct {
	selecting  State
	depositing State
	dispensing State

	current   State
	balance   float64
	basket    []Item
	items     map[string]Item
	sessionID string
	history   map[string]eventstore.Events
}

// NewMachine creates and returns a new vending machine instance.
// It initializes the machine with default values and sets up initial states.
func NewMachine() *Machine {
	machine := &Machine{
		balance: 0,
		items: map[string]Item{ // dummy inventory
			"781": {
				price:    2.5,
				quantity: 3,
			},
			"782": {
				price:    1.5,
				quantity: 5,
			},
		},
		sessionID: uuid.NewString(),
		history:   make(map[string]eventstore.Events),
	}
	selecting := &Selecting{
		machine: machine,
	}
	depositing := &Depositing{
		machine: machine,
	}
	dispensing := &Dispensing{
		machine: machine,
	}

	machine.selecting = selecting
	machine.depositing = depositing
	machine.dispensing = dispensing

	machine.current = selecting
	return machine
}

// Select represents the selecting action on the current state
func (m *Machine) Select(itemID string) error {
	err := m.current.Select(itemID)
	if err != nil {
		return err
	}
	return nil
}

// Deposit represents the depositing action on the current state
func (m *Machine) Deposit(amount float64) error {
	err := m.current.Deposit(amount)
	if err != nil {
		return err
	}
	return nil
}

// Dispense represents the dispensing action on the current state
func (m *Machine) Dispense() error {
	err := m.current.Dispense()
	if err != nil {
		return err
	}
	return nil
}

// Cancel represents the canceling action on the current state
func (m *Machine) Cancel() error {
	err := m.current.Cancel()
	if err != nil {
		return err
	}
	return nil
}

// GetSelectingState getting the Selecting state
func (m *Machine) GetSelectingState() State {
	return m.selecting
}

// GetDepositingState getting the Depositing state
func (m *Machine) GetDepositingState() State {
	return m.depositing
}

// GetDispensingState getting the Dispensing state
func (m *Machine) GetDispensingState() State {
	return m.dispensing
}

// SetState setting a new state
func (m *Machine) SetState(state State) {
	m.current = state
}

// GetTotalPrice compute the total item price in the basket
func (m *Machine) GetTotalPrice() float64 {
	var total float64
	for _, item := range m.basket {
		total += item.price
	}
	return total
}

// GetBalance getting current user balance for the active session
func (m *Machine) GetBalance() float64 {
	return m.balance
}

// GetCurrentState getting current state of the machine
func (m *Machine) GetCurrentState() State {
	return m.current
}

// AddToBalance increase the balance by the given amount
func (m *Machine) AddToBalance(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("unable to deposit below zero: %0.2f\n", amount)
	}
	m.balance += amount
	return nil
}

// SetBalance setting a new balance
func (m *Machine) SetBalance(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("unable to set balance below zero: %0.2f\n", amount)
	}
	m.balance = amount
	return nil
}

// AddToBasket add item to the basket for the active session
func (m *Machine) AddToBasket(itemID string) error {
	item, found := m.items[itemID]
	if !found {
		return fmt.Errorf("unable to find item: %s\n", itemID)
	}
	m.basket = append(m.basket, item)
	return nil
}

// CleanBasket add item to the basket for the active session
func (m *Machine) CleanBasket() {
	m.basket = nil
}

// StoreEvent storing new event to the history
func (m *Machine) StoreEvent(event eventstore.Event) {
	m.history[m.sessionID] = append(m.history[m.sessionID], event)
}
