package engine

import (
	"fmt"
	"github.com/ezzaldeeen/fsvm/eventstore"
)

// Item represents an item in the inventory of the vending engine.
type Item struct {
	name     string
	price    float64
	quantity uint
}

// Engine of the vending engine.
// which is the core component of the vending engine
// where keeps tracking of all states that engine could be in
type Engine struct {
	// predefined states of the engine
	selecting  State
	depositing State
	dispensing State

	currentState State
	balance      float64
	basket       []Item
	inventory    map[string]Item
	history      map[string]eventstore.Events
}

// NewEngine factory function for creating new vending engine's engine
// It initializes the engine with default values and sets up initial states.
func NewEngine() *Engine {
	engine := &Engine{
		balance: 0,
		inventory: map[string]Item{ // dummy inventory
			"1": {
				name:     "pepsi",
				price:    2.5,
				quantity: 3,
			},
			"2": {
				name:     "coke",
				price:    1.5,
				quantity: 5,
			},
		},
		history: make(map[string]eventstore.Events),
	}
	selecting := &Selecting{
		engine: engine,
	}
	depositing := &Depositing{
		engine: engine,
	}
	dispensing := &Dispensing{
		engine: engine,
	}

	engine.selecting = selecting
	engine.depositing = depositing
	engine.dispensing = dispensing

	engine.currentState = selecting
	return engine
}

// Select represents the selecting action on the currentState state
func (e *Engine) Select(itemID string) error {
	err := e.currentState.Select(itemID)
	if err != nil {
		return err
	}
	return nil
}

// Deposit represents the depositing action on the currentState state
func (e *Engine) Deposit(amount float64) error {
	err := e.currentState.Deposit(amount)
	if err != nil {
		return err
	}
	return nil
}

// Dispense represents the dispensing action on the currentState state
func (e *Engine) Dispense() error {
	err := e.currentState.Dispense()
	if err != nil {
		return err
	}
	return nil
}

// Cancel represents the canceling action on the currentState state
func (e *Engine) Cancel() error {
	err := e.currentState.Cancel()
	if err != nil {
		return err
	}
	return nil
}

// CurrentState getting currentState state of the engine
func (e *Engine) CurrentState() State {
	return e.currentState
}

// SetState setting a new state
func (e *Engine) setState(state State) {
	e.currentState = state
}

// getTotalPrice compute the total item price in the currentBasket
func (e *Engine) getTotalPrice() float64 {
	var total float64
	for _, item := range e.basket {
		total += item.price
	}
	return total
}

// addToBalance increase the balance by the given amount
func (e *Engine) addToBalance(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf(
			"unable to deposit below zero: %0.2f", amount,
		)
	}
	e.balance += amount
	return nil
}

// addToBasket add item to the currentBasket for the active session
func (e *Engine) addToBasket(itemID string) error {
	item, found := e.inventory[itemID]
	if !found {
		return fmt.Errorf(
			"item with id: %s, not in the inventory", itemID,
		)
	}
	e.basket = append(e.basket, item)
	return nil
}

// reset current balance and current basket
func (e *Engine) reset() {
	e.balance = 0
	e.basket = nil
}
