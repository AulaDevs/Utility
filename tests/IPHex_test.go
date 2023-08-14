package tests

import (
	"testing"
	
	. "github.com/AulaDevs/Utility"
)

func TestIPHexHandler(t *testing.T) {
	ips := []string{
		"127.0.0.1", 
		"23.67.11.129", 
		"231.45.74.1", 
		"62.217.12.74", 
		"102.35.67.211",
		"128.128.128.11",
		"69.42.6.66", // 69|42|666
		"12.14.16.18",
	}
	
	results := []string{
		"#7F.00.00.01",
		"#17.43.0B.81", 
		"#E7.2D.4A.01", 
		"#3E.D9.0C.4A", 
		"#66.23.43.D3",
		"#80.80.80.0B",
		"#45.2A.06.42",
		"#0C.0E.10.12",
	}

	for i, str := range ips {
		str = EncodeIP(str)
		if str != results[i] {
			t.Fatalf("Invalid IP Encoding. Expected: %s | Got: %s", results[i], str)
			break
		}
	}

	t.Logf("Test #1 Passed")
	
	for i, str := range results {
		str = DecodeIP(str)
		if str != ips[i] {
			t.Fatalf("Invalid IP Decoding. Expected: %s | Got: %s", ips[i], str)
			break
		}
	}
	
	t.Logf("Test #2 Passed. All tests are passed.")
}
