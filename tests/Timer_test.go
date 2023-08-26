package tests

import (
	"testing"

	"github.com/AulaDevs/Utility/timer"
)

func TestTimer(t *testing.T) {
	canExit := make(chan bool)

	timer.CallLater(
		2,
		func(args []any) {
			t.Logf("Args passed after 2s: %v", args)
			canExit <- true
		},
		"Hello world",
	)

	<-canExit
}

func TestTimerCancel(t *testing.T) {
	tm := timer.CallLater(
		2,
		func(args []any) {
			t.Logf("Args passed after 2s: %v", args) // it never run
			t.Fail()
		},
		"Hello world",
	)

	t.Log("Cancelling goroutine")
	if err := tm.Cancel(); err != nil {
		t.Fatal(err)
	}

	for !tm.IsFinished() {

	}
}
