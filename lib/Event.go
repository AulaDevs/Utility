package lib

import "github.com/Goldziher/go-utils/sliceutils"

type Event []interface{}

type EventListener struct {
	once     bool
	callback func(Event)
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
func (eventHandler *EventHandler) ListenOnce(name string, callback func(Event)) {
	eventHandler.listeners[name] = append(
		eventHandler.listeners[name],
		&EventListener{
			once:     true,
			callback: callback,
		},
	)
}

func (eventHandler *EventHandler) Listen(name string, callback func(Event)) {
	eventHandler.listeners[name] = append(
		eventHandler.listeners[name],
		&EventListener{
			once:     false,
			callback: callback,
		},
	)
}

func (eventHandler *EventHandler) ListenAndWait(name string) Event {
	channel := make(chan Event)

	eventHandler.ListenOnce(
		name,
		func(args Event) {
			channel <- args
		},
	)

	return <-channel
}

func (eventHandler *EventHandler) Emit(name string, data Event) {
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
