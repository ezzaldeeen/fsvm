package machine

// State is an interface that represents
// the state of a vending machine.
type State interface {
	Select(itemID string) error
	Deposit(amount float64) error
	Dispense() error
	Cancel() error
}
