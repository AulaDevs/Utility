package net

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type Socket struct {
	conn   *net.TCPConn
	closed bool
}

// Constructors
func SocketFromTCPConn(conn *net.TCPConn) *Socket {
	return &Socket{conn, false}
}

func Connect(host string, port int) (*Socket, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}

	return &Socket{conn.(*net.TCPConn), false}, nil
}

// Standard methods
func (socket *Socket) IsClosed() bool {
	return socket.closed
}

func (socket *Socket) Refresh() Response {
	if socket.closed {
		return Response{NET_CLOSED, []any{}}
	}

	socket.conn.SetDeadline(time.Now().Add(2 * time.Millisecond))

	data := make([]byte, 2048)
	bytesRead, err := socket.conn.Read(data)

	if err != nil {
		if os.IsTimeout(err) {
			return Response{NET_SKIP, []any{}}
		} else if err == io.EOF {
			socket.Close()
			return Response{NET_CLOSED, []any{}}
		}

		return Response{NET_ERR, []any{err}}
	}

	if bytesRead != 0 {
		buffer := bytes.NewBuffer(make([]byte, 0, bytesRead))
		buffer.Write(data[:bytesRead])

		return Response{NET_DATA, []any{buffer}}
	}

	return Response{NET_SKIP, []any{}}
}

func (socket *Socket) Close() {
	socket.closed = true
	socket.conn.Close()
}

func (socket *Socket) Write(buffer *bytes.Buffer) (int, error) {
	if socket.closed {
		return 0, nil
	}

	return socket.conn.Write(buffer.Bytes())
}

func (socket *Socket) RemoteAddress() net.Addr {
	return socket.conn.RemoteAddr()
}

func (socket *Socket) LocalAddress() net.Addr {
	return socket.conn.LocalAddr()
}
