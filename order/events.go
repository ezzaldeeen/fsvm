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
	// Pending ...
	Pending struct {
		orderID string
	}

	// Selected ...
	Selected struct {
		orderID   string
		productID string
	}

	// Deposited ...
	Deposited struct {
		orderID string
		amount  uint
	}

	// Canceled ...
	Canceled struct {
		orderID string
	}

	// Purchased ...
	Purchased struct {
		orderID string
	}
)
