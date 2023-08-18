package event

import "github.com/Goldziher/go-utils/sliceutils"

type Args []interface{}

type EventListener struct {
	once     bool
	callback func(Args)
}

type EventHandler struct {
	listeners map[string][]*EventListener
}

// Constructor
func NewEventHandler() *EventHandler {
	return &EventHandler{
		listeners: map[string][]*EventListener{},
	}
}

// Standard methods
func (eventHandler *EventHandler) ListenOnce(name string, callback func(Args)) {
	eventHandler.listeners[name] = append(
		eventHandler.listeners[name],
		&EventListener{
			once:     true,
			callback: callback,
		},
	)
}

func (eventHandler *EventHandler) Listen(name string, callback func(Args)) {
	eventHandler.listeners[name] = append(
		eventHandler.listeners[name],
		&EventListener{
			once:     false,
			callback: callback,
		},
	)
}

func (eventHandler *EventHandler) ListenAndWait(name string) Args {
	channel := make(chan Args)

	eventHandler.ListenOnce(
		name,
		func(args Args) {
			channel <- args
		},
	)

	return <-channel
}

func (eventHandler *EventHandler) Emit(name string, data Args) {
	if len(eventHandler.listeners[name]) == 0 {
		return
	}

	for index, listener := range eventHandler.listeners[name] {
		go func() {
			listener.callback(data)
		}()

		if listener.once {
			eventHandler.listeners[name] = sliceutils.Remove(eventHandler.listeners[name], index)
		}
	}
}

func (eventHandler *EventHandler) CountEventListeners(name string) int {
	return len(eventHandler.listeners[name])
}

func (eventHandler *EventHandler) RemoveAllEventListeners() {
	eventHandler.listeners = map[string][]*EventListener{}
}

func (eventHandler *EventHandler) RemoveEventListeners(name string) {
	eventHandler.listeners[name] = nil
}
