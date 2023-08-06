package tests

import (
	"bytes"
	"testing"

	. "github.com/AulaDevs/Utility/lib"
)

func TestNetworkListen(t *testing.T) {
	network, err := NetworkListen("0.0.0.0", 11801)

	if err != nil {
		t.Fatal(err)
	}

	network.ListenEvent(
		"Error",
		func(args Event) {
			t.Fatalf("%s: %v", args[0].(string), args[1].(error))
		},
	)

	network.ListenEvent(
		"NewClient",
		func(args Event) {
			socket := args[0].(*NetworkSocket)

			t.Logf("NewClient %s", socket.RemoteAddress().String())

			socket.ListenEvent(
				"Error",
				func(args Event) {
					t.Fatalf("%s: %v", args[0].(string), args[1].(error))
				},
			)

			socket.ListenEvent(
				"DataReceived",
				func(args Event) {
					buffer := args[0].(*bytes.Buffer)

					t.Logf("Received bytes: %v", buffer.Bytes())

					socket.Close()
				},
			)

			socket.ListenEvent("Closed", func(_ Event) {
				t.Log("Socket closed")
				network.Close()
			})
		},
	)

	canExit := make(chan bool)
	network.ListenEvent("Closed", func(_ Event) {
		t.Log("Network closed")
		canExit <- true
	})

	socket, err := NetworkSocketConnect("0.0.0.0", 11801)

	if err != nil {
		t.Fatal(err)
	}

	sent, err := socket.Write(bytes.NewBufferString("Hello world"))

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Socket sent %d bytes to the network socket", sent)

	<-canExit
}
