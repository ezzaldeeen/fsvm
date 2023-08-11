package main

import (
	"github.com/ezzaldeeen/fsvm/engine"
	"github.com/ezzaldeeen/fsvm/eventstream"
	"log"
	"sync"
)

const numOfGoroutines = 2

func main() {
	wg := new(sync.WaitGroup)
	// initialize vending engine
	eng := engine.NewEngine()
	// channel for event propagation
	events := make(chan eventstream.Event)
	// dispatcher for the occurred events
	dispatcher := eventstream.NewEventDispatcher(eng, events)
	// handler for handle the dispatched events
	handler := eventstream.NewEventHandler(eng, events)
	// running the vending engine
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
