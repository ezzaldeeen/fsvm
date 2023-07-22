package order

import "github.com/google/uuid"

// State represents the order's state
type State string

// History chain of ordered events
type History []Event

const (
	PENDING   State = "PENDING"
	SELECTED  State = "SELECTED"
	DEPOSITED State = "DEPOSITED"
	CANCELED  State = "CANCELED"
	PURCHASED State = "PURCHASED"
)

// Order represents the session in vending machine,
// where each session has only one order
type Order struct {
	id        string
	productID string
	state     State
	balance   uint
	history   History
	version   uint
}

// NewOrder create new order with PENDING (initial) event
func NewOrder() *Order {
	o := &Order{}
	o.raise(Pending{
		orderID: uuid.NewString(),
	})
	return o
}

// Replay rewind the final state from all previous states
func Replay(history History) *Order {
	o := &Order{}
	for _, event := range history {
		o.on(event, false)
	}
	return o
}

// GetID getting the id of the order
func (o *Order) GetID() string {
	return o.id
}

// GetState getting the state of the order
func (o *Order) GetState() State {
	return o.state
}

// GetHistory getting the event history of the order
func (o *Order) GetHistory() History {
	return o.history
}

// GetVersion getting the version of the order
// used for optimistic concurrency
func (o *Order) GetVersion() uint {
	return o.version
}

// on ..
func (o *Order) on(event Event, new bool) {
	switch e := event.(type) {
	case Pending:
		o.id = e.orderID
		o.state = PENDING
	case Selected:
		o.id = e.orderID
		o.productID = e.productID
		o.state = SELECTED
	case Canceled:
		o.id = e.orderID
		o.state = CANCELED
	case Deposited:
		o.id = e.orderID
		o.balance += e.amount
		o.state = DEPOSITED
	case Purchased:
		o.id = e.orderID
		o.state = PURCHASED
	}
	if !new {
		o.version++
	}
}

// raise new event over the chain (History)
func (o *Order) raise(event Event) {
	o.history = append(o.history, event)
	o.on(event, true)
}
