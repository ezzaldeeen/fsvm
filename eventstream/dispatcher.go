package eventstream

import (
	"github.com/ezzaldeeen/fsvm/machine"
	"time"
)

// EventDispatcher is responsible for dispatching events to
// a channel from the user who interact with the vending machine.
type EventDispatcher struct {
	machine *machine.Machine
	events  chan<- Event
}

// NewEventDispatcher creates a new EventDispatcher instance
func NewEventDispatcher(machine *machine.Machine,
	events chan<- Event) *EventDispatcher {
	return &EventDispatcher{
		events:  events,
		machine: machine,
	}
}

// Run starts the event dispatcher to monitor the state of a vending machine
// and send corresponding events when certain state transitions (event) occur.
func (d EventDispatcher) Run() {
	defer close(d.events)
	for {
		time.Sleep(10 * time.Millisecond)
		switch d.machine.GetCurrentState().(type) {
		case *machine.Selecting:
			itemID := machine.DoSelectingProcess()
			event := &Selected{ItemID: itemID}
			d.events <- event
		case *machine.Depositing:
			amount := machine.DoDepositingProcess()
			event := &Deposited{Amount: amount}
			d.events <- event
		case *machine.Dispensing:
			event := &Dispensed{}
			d.events <- event
		}
	}
}
