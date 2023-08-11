package engine

// State is an interface that represents
// the state of a vending engine.
type State interface {
	Select(itemID string) error
	Deposit(amount float64) error
	Dispense() error
	Cancel() error
}
