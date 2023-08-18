package net

import (
	"bytes"
	"fmt"
	"net"

	"github.com/AulaDevs/Utility/event"
)

type NetworkSocket struct {
	conn   *net.TCPConn
	Events *event.EventHandler
	closed bool
}

// Constructors
func NetworkSocketFrom(conn *net.TCPConn) *NetworkSocket {
	return &NetworkSocket{conn, event.NewEventHandler(), false}
}

func NetworkSocketConnect(host string, port int) (*NetworkSocket, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}

	return &NetworkSocket{conn.(*net.TCPConn), event.NewEventHandler(), false}, nil
}

// Standard methods

func (socket *NetworkSocket) Poll() {
	go func() {
		defer socket.Close()

		for !socket.closed {
			data := make([]byte, 2048)
			bytesRead, err := socket.conn.Read(data)

			if socket.closed {
				return
			}

			if err != nil {
				socket.Events.Emit("Error", event.Args{"DataReceived", err})
				return
			}

			if bytesRead != 0 {
				buffer := bytes.NewBuffer(make([]byte, 0, bytesRead))
				buffer.Write(data[:bytesRead])

				socket.Events.Emit("DataReceived", event.Args{buffer})
			}
		}
	}()
}

func (socket *NetworkSocket) Close() {
	socket.closed = true
	socket.conn.Close()
	socket.Events.Emit("Closed", event.Args{})
	socket.Events.RemoveAllEventListeners()
}

func (socket *NetworkSocket) Write(buffer *bytes.Buffer) (int, error) {
	if socket.closed {
		return 0, nil
	}

	return socket.conn.Write(buffer.Bytes())
}

func (socket *NetworkSocket) RemoteAddress() net.Addr {
	return socket.conn.RemoteAddr()
}

func (socket *NetworkSocket) LocalAddress() net.Addr {
	return socket.conn.LocalAddr()
}

// Event methods
func (socket *NetworkSocket) ListenEvent(name string, callback func(event.Args)) {
	socket.Events.Listen(name, callback)
}

func (socket *NetworkSocket) ListenEventOnce(name string, callback func(event.Args)) {
	socket.Events.ListenOnce(name, callback)
}
