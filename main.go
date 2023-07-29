package main

import (
	"github.com/ezzaldeeen/fsvm/eventstream"
	"github.com/ezzaldeeen/fsvm/machine"
	"log"
	"sync"
)

const numOfGoroutines = 2

func main() {
	wg := new(sync.WaitGroup)
	// initialize vending machine
	mac := machine.NewMachine()
	// channel for event propagation
	events := make(chan eventstream.Event)
	// dispatcher for the occurred events
	dispatcher := eventstream.NewEventDispatcher(mac, events)
	// handler for handle the dispatched events
	handler := eventstream.NewEventHandler(mac, events)

	wg.Add(numOfGoroutines)
	go dispatcher.Run()
	go func() {
		err := handler.Handle()
		if err != nil {
			log.Println(err)
		}
	}()
	wg.Wait()
}
