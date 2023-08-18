package tests

import (
	"testing"

	. "github.com/AulaDevs/Utility/event"
)

func TestEventHandler(t *testing.T) {
	canExit := make(chan bool)
	canProceed := make(chan bool)

	var handler *EventHandler = NewEventHandler()

	go func() {
		canProceed <- true
		args := handler.ListenAndWait("test")
		t.Logf("Event 'test' called with args: %v", args)
		canExit <- true
	}()

	<-canProceed

	t.Log("Running 'test' Emitter")
	handler.Emit("test", Args{3, "Lucas"})

	<-canExit

	if handler.CountEventListeners("test") > 0 {
		t.Fatal("It was expected that there would be no listeners registered at the end of the test.")
	}
}
