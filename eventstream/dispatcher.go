package eventstream

import (
	"fmt"
	"github.com/ezzaldeeen/fsvm/engine"
	"strconv"
	"time"
)

// EventDispatcher is responsible for dispatching events to
// a channel from the user who interact with the vending engine.
type EventDispatcher struct {
	engine *engine.Engine
	events chan<- Event
}

// NewEventDispatcher creates a new EventDispatcher instance
func NewEventDispatcher(engine *engine.Engine,
	events chan<- Event) *EventDispatcher {
	return &EventDispatcher{
		events: events,
		engine: engine,
	}
}

// Run starts the event dispatcher to monitor the state of a vending engine
// and send corresponding events when certain state transitions (event) occur.
func (d EventDispatcher) Run() {
	defer close(d.events)
	for {
		// to make sure that console's output consistent
		time.Sleep(time.Second)
		switch d.engine.CurrentState().(type) {
		case *engine.Selecting:
			fmt.Print("[SELECTING] - please enter item code > ")
			var itemID string
			_, err := fmt.Scan(&itemID)
			if err != nil {
				fmt.Println(err)
			}
			// propagate event
			event := &Selected{ItemID: itemID}
			d.events <- event

		case *engine.Depositing:
			fmt.Print("[DEPOSITING] - enter money > ")
			var input string
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Println(err)
			}
			// in case the user want to cancel the order
			if input == ":q" {
				// propagate event
				event := &Canceled{}
				d.events <- event
			}
			amount, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("[DEPOSITING] - invalid input. expected to get: 1.00, 2.5, ...")
			}
			// propagate event
			event := &Deposited{Amount: amount}
			d.events <- event

		case *engine.Dispensing:
			fmt.Println("[DISPENSING] - dispensing in progress...")
			// propagate event
			event := &Dispensed{}
			d.events <- event
		}
	}
}
