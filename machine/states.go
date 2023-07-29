package machine

type State interface {
	Select(itemID string) error
	Deposit(amount float64) error
	Dispense() error
	Cancel() error
}
