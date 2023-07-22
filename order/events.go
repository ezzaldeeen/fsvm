package order

type Event interface {
	event()
}

func (e Pending) event()   {}
func (e Selected) event()  {}
func (e Deposited) event() {}
func (e Canceled) event()  {}
func (e Purchased) event() {}

type (
	// Pending event when the user start interacting
	// with the vending machine, i.e. the initial event
	Pending struct {
		orderID string
	}

	// Selected event when the product selected in that order
	Selected struct {
		orderID   string
		productID string
	}

	// Deposited event when the money for that order deposited
	Deposited struct {
		orderID string
		amount  uint
	}

	// Canceled event when the order is canceled
	Canceled struct {
		orderID string
	}

	// Purchased event when the order is completed
	Purchased struct {
		orderID string
	}
)
