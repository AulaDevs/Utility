package tests

import (
	"bytes"
	crand "crypto/rand"
	"math/rand"
	"testing"
	"time"

	. "github.com/AulaDevs/Utility/lib"
)

var rng *rand.Rand
var byteArray *ByteArray = NewByteArrayEmpty()

func CheckCreateRandomSeed(t *testing.T) {
	if rng == nil {
		var seed = rand.Int63n(time.Now().UnixNano()) // need to randomize it
		rng = rand.New(rand.NewSource(seed))
		t.Logf("Using seed %d", seed)
	}
}

func TestByteArrayByte(t *testing.T) {
	CheckCreateRandomSeed(t)

	// Up to max byte value
	randomByte := byte(rng.Intn(256))

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
	CheckCreateRandomSeed(t)

	// Up to max short value
	randomShort := uint16(rng.Intn(65536))

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
	CheckCreateRandomSeed(t)

	// Up to max int value
	randomInt := rng.Intn(2147483648)

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

func TestByteArrayInt48(t *testing.T) {
	CheckCreateRandomSeed(t)

	// Up to max int64 value
	randomDynamicInt := rng.Int63n(281474976710655)

	byteArray.Write_Int48(randomDynamicInt)

	t.Logf("Wrote number %d, bytes: %v", randomDynamicInt, byteArray.Get_Bytes())

	integerRead := byteArray.Read_Int48()

	if integerRead != randomDynamicInt {
		t.Fatalf("The dynamic int read was expected to be %d but got %d.", randomDynamicInt, integerRead)
	}

	if byteArray.Len() > 0 {
		t.Fatalf("It was expected that after reading the written bytes the buffer would be empty, but there are still %d bytes in the buffer. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}

	byteArray.Clear()
}

// For testing string purposes
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}

func TestByteArrayString(t *testing.T) {
	CheckCreateRandomSeed(t)

	randomSize := rng.Intn(65536)
	randomString := RandomString(randomSize)

	fullSize := 2 + randomSize // short + string bytes

	byteArray.Write_String(randomString)

	if byteArray.Len() != fullSize {
		t.Fatalf("It was expected that after writing string the buffer size would be %d but it is %d. Bytes: %v", fullSize, byteArray.Len(), byteArray.Get_Bytes())
	}

	stringRead := byteArray.Read_String()

	if stringRead != randomString {
		t.Fatalf("The string read was expected to be %s but got %s.", randomString, stringRead)
	}

	if byteArray.Len() > 0 {
		t.Fatalf("It was expected that after reading the written bytes the buffer would be empty, but there are still %d bytes in the buffer. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}
}

func TestByteArrayBool(t *testing.T) {
	CheckCreateRandomSeed(t)

	randomBool := rng.Intn(2) == 1

	byteArray.Write_Boolean(randomBool)

	if byteArray.Len() != 1 {
		t.Fatalf("It was expected that after writing boolean the buffer size would be 1 but it is %d. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}

	boolRead := byteArray.Read_Boolean()

	if boolRead != randomBool {
		t.Fatalf("The boolean read was expected to be %v but got %v.", randomBool, boolRead)
	}

	if byteArray.Len() > 0 {
		t.Fatalf("It was expected that after reading the written bytes the buffer would be empty, but there are still %d bytes in the buffer. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}
}

func TestByteArrayBytes(t *testing.T) {
	CheckCreateRandomSeed(t)

	randomBytes := make([]byte, rng.Intn(64))
	randomSize, err := crand.Read(randomBytes)

	if err != nil {
		t.Fatal(err)
	}

	byteArray.Write_Bytes(randomBytes)

	if byteArray.Len() != randomSize {
		t.Fatalf("It was expected that after writing bytes the buffer size would be %d but it is %d. Bytes: %v", randomSize, byteArray.Len(), byteArray.Get_Bytes())
	}

	bytesRead := byteArray.Read_Bytes(randomSize)

	if !bytes.Equal(bytesRead, randomBytes) {
		t.Fatalf("The bytes read was expected to be %v but got %v.", randomBytes, bytesRead)
	}

	if byteArray.Len() > 0 {
		t.Fatalf("It was expected that after reading the written bytes the buffer would be empty, but there are still %d bytes in the buffer. Bytes: %v", byteArray.Len(), byteArray.Get_Bytes())
	}
}
