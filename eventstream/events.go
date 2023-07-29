package eventstream

// Event represents the event that occurred by user
type Event interface {
	event()
}

// event workaround to make these events
// implements the Event interface
func (s Selected) event()  {}
func (s Deposited) event() {}
func (s Dispensed) event() {}
func (s Canceled) event()  {}

// Selected event when the user selected an item
type Selected struct {
	ItemID string
}

// Deposited event when the user deposited a money
type Deposited struct {
	Amount float64
}

// Dispensed event when the user dispensed the item, and remaining money
type Dispensed struct {
}

// Canceled event when the user decided to cancel purchase
type Canceled struct {
}
