package Utility

import (
	"fmt"
	"net"
)

type Network struct {
	socket *net.TCPListener
	Events *EventHandler
	closed bool
}

// Constructor
func NetworkListen(host string, port int) (*Network, error) {
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		return nil, err
	}

	network := &Network{
		socket: listener.(*net.TCPListener),
		Events: NewEventHandler(),
		closed: false,
	}

	go func() {
		for !network.closed {
			conn, err := network.socket.AcceptTCP()

			if err != nil {
				network.Events.Emit("Error", Event{"NewClient", err})
			} else {
				network.Events.Emit("NewClient", Event{NetworkSocketFrom(conn)})
			}
		}
	}()

	return network, nil
}

func (network *Network) Close() {
	network.closed = true
	network.socket.Close()
	network.Events.Emit("Closed", Event{})
	network.Events.RemoveAllEventListeners()
}

// Event methods
func (network *Network) ListenEvent(name string, callback func(Event)) {
	network.Events.Listen(name, callback)
}

func (network *Network) ListenEventOnce(name string, callback func(Event)) {
	network.Events.ListenOnce(name, callback)
}
