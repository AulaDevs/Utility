package timer

import (
	"errors"
	"time"
)

type Timer struct {
	canceled bool
	runned   bool
	cancel   chan bool
}

func CallLater(seconds int, callback func(args []any), args ...any) *Timer { //can u let me do it then u can do? sure
	timer := &Timer{
		canceled: false,
		runned:   false,
		cancel:   make(chan bool),
	}

	go func() {
		timeout := time.NewTimer(time.Duration(seconds) * time.Second)

		select {
		case <-timer.cancel:
			if timer.runned {
				return
			}

			timer.canceled = true

			if !timeout.Stop() {
				<-timeout.C
			}

			close(timer.cancel)
		case <-timeout.C:
			if timer.canceled {
				return
			}

			callback(args)
			timer.runned = true
			timer.cancel <- true
		}
	}()

	return timer
}

func (timer *Timer) IsCanceled() bool {
	return timer.canceled
}

func (timer *Timer) IsFinished() bool {
	return timer.runned || timer.canceled
}

func (timer *Timer) Cancel() error {
	if timer.runned {
		return errors.New("Timer already runned")
	}

	if timer.canceled {
		return errors.New("Timer already canceled")
	}

	timer.cancel <- true

	return nil
}
