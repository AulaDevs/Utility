package lib

import (
	"bytes"
	"fmt"
	"net"
)

type NetworkSocket struct {
	conn   *net.TCPConn
	Events *EventHandler
	closed bool
}

// Constructors
func NetworkSocketFrom(conn *net.TCPConn) *NetworkSocket {
	socket := &NetworkSocket{conn, NewEventHandler(), false}
	go socket.pollData()
	return socket
}

func NetworkSocketConnect(host string, port int) (*NetworkSocket, error) {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}

	socket := &NetworkSocket{conn.(*net.TCPConn), NewEventHandler(), false}
	go socket.pollData()

	return socket, nil
}

// Private methods
func (socket *NetworkSocket) pollData() {
	defer socket.Close()

	for !socket.closed {
		data := make([]byte, 2048)
		bytesRead, err := socket.conn.Read(data)

		if socket.closed {
			return
		}

		if err != nil {
			socket.Events.Emit("Error", Event{"DataReceived", err})
			return
		}

		if bytesRead != 0 {
			buffer := bytes.NewBuffer(make([]byte, 0, bytesRead))
			buffer.Write(data[:bytesRead])

			socket.Events.Emit("DataReceived", Event{buffer})
		}
	}
}

// Standard methods
func (socket *NetworkSocket) Close() {
	socket.closed = true
	socket.conn.Close()
	socket.Events.Emit("Closed", Event{})
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
func (socket *NetworkSocket) ListenEvent(name string, callback func(Event)) {
	socket.Events.Listen(name, callback)
}

func (socket *NetworkSocket) ListenEventOnce(name string, callback func(Event)) {
	socket.Events.ListenOnce(name, callback)
}
