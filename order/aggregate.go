package order

type State string
type History []Event

const (
	PENDING   State = "PENDING"
	SELECTED  State = "SELECTED"
	DEPOSITED State = "DEPOSITED"
	CANCELED  State = "CANCELED"
	PURCHASED State = "PURCHASED"
)

// Order ...
type Order struct {
	id        string // todo: use uuid
	productID string
	state     State
	balance   uint
	history   History
	version   uint
}

func NewOrder() *Order {
	return &Order{
		id:      "", // todo: use uuid
		state:   "", // todo: set default state (initial)
		balance: 0,
		history: nil,
		version: 0,
	}
}

// Replay ...
func Replay(history History) *Order {
	return nil // todo: implement
}

// GetID ID ...
func (o *Order) GetID() string {
	return o.id
}

// GetState State ...
func (o *Order) GetState() State {
	return o.state
}

// GetHistory ...
func (o *Order) GetHistory() History {
	return o.history
}

// GetVersion Version ...
func (o *Order) GetVersion() uint {
	return o.version
}

// on ..
func (o *Order) on(event Event) { // todo: implement versioning
	switch e := event.(type) {
	case Pending:
		o.id = e.orderID
		o.state = SELECTED
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
}

// raise ..
func (o *Order) raise(event Event) {
	o.history = append(o.history, event)
	o.on(event)
}
