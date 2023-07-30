package tests

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/AulaDevs/Utility/lib"
)

var byteArray *ByteArray = NewByteArrayEmpty()

func TestByteArrayByte(t *testing.T) {
	// Up to max byte value
	randomByte := byte(rand.Intn(255))

	byteArray.Write_Byte(randomByte)

	if byteArray.Len() != 1 {
		t.Fatalf("It was expected that after writing byte the buffer size would be 1 but it is %d. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}

	byteRead := byteArray.Read_Byte()

	if byteRead != randomByte {
		t.Fatalf("The byte read was expected to be %d but got %d.", randomByte, byteRead)
	}

	if byteArray.Len() > 0 {
		t.Fatalf("It was expected that after reading the written bytes the buffer would be empty, but there are still %d bytes in the buffer. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}
}

func TestByteArrayShort(t *testing.T) {
	// Up to max short value
	randomShort := uint16(rand.Intn(65535))

	byteArray.Write_Short(randomShort)

	if byteArray.Len() != 2 {
		t.Fatalf("It was expected that after writing short the buffer size would be 2 but it is %d. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}

	shortRead := byteArray.Read_Short()

	if shortRead != randomShort {
		t.Fatalf("The short read was expected to be %d but got %d.", randomShort, shortRead)
	}

	if byteArray.Len() > 0 {
		t.Fatalf("It was expected that after reading the written bytes the buffer would be empty, but there are still %d bytes in the buffer. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}
}

func TestByteArrayInt(t *testing.T) {
	// Up to max int value
	randomInt := rand.Intn(2147483647)

	byteArray.Write_Int(randomInt)

	if byteArray.Len() != 4 {
		t.Fatalf("It was expected that after writing int the buffer size would be 4 but it is %d. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}

	integerRead := byteArray.Read_Int()

	if integerRead != randomInt {
		t.Fatalf("The int read was expected to be %d but got %d.", randomInt, integerRead)
	}

	if byteArray.Len() > 0 {
		t.Fatalf("It was expected that after reading the written bytes the buffer would be empty, but there are still %d bytes in the buffer. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}
}

// For testing string purposes
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func TestByteArrayString(t *testing.T) {
	randomSize := rand.Intn(65535)
	randomString := RandomString(randomSize)

	fullSize := 2 + randomSize // short + string bytes

	byteArray.Write_String(randomString)

	if byteArray.Len() != fullSize {
		t.Fatalf("It was expected that after writing int the buffer size would be %d but it is %d. Bytes: %v", fullSize, byteArray.Len(), byteArray.Get_Bytes())
	}

	stringRead := byteArray.Read_String()

	if stringRead != randomString {
		t.Fatalf("The string read was expected to be %s but got %s.", randomString, stringRead)
	}

	if byteArray.Len() > 0 {
		t.Fatalf("It was expected that after reading the written bytes the buffer would be empty, but there are still %d bytes in the buffer. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}
}
