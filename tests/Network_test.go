package tests

import (
	"bytes"
	"testing"

	"github.com/AulaDevs/Utility/net"
)

func handleClient(t *testing.T, socket *net.Socket, network *net.Listener) {
	t.Logf("New socket from: %s", socket.RemoteAddress())
	for !socket.IsClosed() {
		response := socket.Refresh()

		switch response.Type {
		case net.NET_CLOSED:
			t.Log("Socket closed")
			network.Close()
		case net.NET_ERR:
			t.Logf("Error: %v", response.Data[0].(error))
		case net.NET_DATA:
			t.Logf("Received bytes: %v", response.Data[0].(*bytes.Buffer).String())
		case net.NET_SKIP:
			continue
		}
	}
}

func TestNetworkListen(t *testing.T) {
	canExit := make(chan bool)

	network, err := net.Listen("0.0.0.0", 11801)

	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for !network.IsClosed() {
			response := network.Refresh()

			switch response.Type {
			case net.NET_CLOSED:
				t.Log("Network closed")
			case net.NET_CONN:
				go handleClient(t, response.Data[0].(*net.Socket), network)
			case net.NET_ERR:
				t.Logf("Error: %v", response.Data[0].(error))
			case net.NET_SKIP:
				continue
			}
		}

		canExit <- true
	}()

	client, err := net.Connect("0.0.0.0", 11801)

	if err != nil {
		t.Fatal(err)
	}

	client.Write(bytes.NewBufferString("Hello world"))
	client.Close()

	<-canExit
}
