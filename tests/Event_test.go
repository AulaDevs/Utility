package tests

import (
	"testing"

	. "github.com/AulaDevs/Utility/lib"
)

var handler *EventHandler = NewEventHandler()

func TestEventHandler(t *testing.T) {
	canExit := make(chan bool)

	go func() {
		args := handler.ListenAndWait("test")
		t.Logf("Event 'test' called with args: %v", args)
		canExit <- true
	}()

	t.Log("Running 'test' Emitter")
	handler.Emit("test", Event{3, "Lucas"})

	<-canExit

	if handler.CountEventListeners("test") > 0 {
		t.Fatal("It was expected that there would be no listeners registered at the end of the test.")
	}
}
