package eventstream

import (
	"fmt"
	"github.com/ezzaldeeen/fsvm/machine"
)

// EventHandler represents an event handler for a vending machine.
// It listens for events from the channel and processes them accordingly.
type EventHandler struct {
	machine *machine.Machine
	events  <-chan Event
}

// NewEventHandler creates and returns a new event handler instance.
func NewEventHandler(machine *machine.Machine,
	events <-chan Event) *EventHandler {
	return &EventHandler{
		machine: machine,
		events:  events,
	}
}

// Handle listens for events from the channel and processes them accordingly.
// It handles different types of events and takes appropriate actions
// based on the event type. It calls methods on the associated vending machine
// to perform actions like selecting, depositing, dispensing, or canceling.
// The function returns nil if all events are processed successfully.
func (h EventHandler) Handle() error {
	for event := range h.events {
		switch e := event.(type) {
		case *Selected:
			err := h.machine.Select(e.ItemID)
			if err != nil {
				fmt.Println(err)
			}
		case *Deposited:
			err := h.machine.Deposit(e.Amount)
			if err != nil {
				fmt.Println(err)
			}
		case *Dispensed:
			err := h.machine.Dispense()
			if err != nil {
				fmt.Println(err)
			}
		case *Canceled:
			err := h.machine.Cancel()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
