package net

import (
	"fmt"
	"net"
	"os"
	"time"
)

type Listener struct {
	socket *net.TCPListener
	closed bool
}

// Response types
const (
	NET_SKIP = iota
	NET_CONN
	NET_DATA
	NET_CLOSED
	NET_ERR
)

type Response struct {
	Type int
	Data []any
}

// Constructor
func Listen(host string, port int) (*Listener, error) {
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		return nil, err
	}

	return &Listener{
		socket: listener.(*net.TCPListener),
		closed: false,
	}, nil
}

func (network *Listener) IsClosed() bool {
	return network.closed
}

func (network *Listener) Refresh() Response {
	if network.closed {
		return Response{NET_CLOSED, []any{}}
	}

	network.socket.SetDeadline(time.Now().Add(2 * time.Millisecond))

	conn, err := network.socket.AcceptTCP()

	if err != nil {
		if os.IsTimeout(err) {
			return Response{NET_SKIP, []any{}}
		}

		return Response{NET_ERR, []any{err}}
	}

	return Response{NET_CONN, []any{SocketFromTCPConn(conn)}}
}

func (network *Listener) Close() {
	network.closed = true
	network.socket.Close()
}
