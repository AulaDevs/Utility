package tests

import (
	"bytes"
	"math"
	"math/rand"
	"testing"

	. "github.com/AulaDevs/Utility"
)

func TestNetworkListen(t *testing.T) {
	port := int(math.Min(10000, float64(rand.Intn(160000))))

	network, err := NetworkListen("0.0.0.0", port)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Listening on 0.0.0.0:%d", port)

	network.ListenEvent(
		"Error",
		func(args Event) {
			t.Logf("%s: %v", args[0].(string), args[1].(error))
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

			socket.Poll()
		},
	)

	canExit := make(chan bool)
	network.ListenEvent("Closed", func(_ Event) {
		t.Log("Network closed")
		canExit <- true
	})

	socket, err := NetworkSocketConnect("0.0.0.0", port)

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
