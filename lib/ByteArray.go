package lib

import (
	"bytes"
	"encoding/binary"
)

type ByteArray struct {
	Bytes *bytes.Buffer
}

// Constructors
func NewByteArrayEmpty() *ByteArray {
	return &ByteArray{Bytes: bytes.NewBuffer([]byte{})}
}

func NewByteArray(bytes *bytes.Buffer) *ByteArray {
	return &ByteArray{Bytes: bytes}
}

// Standard methods
func (bt *ByteArray) Clear() *ByteArray {
	bt.Bytes.Reset()
	return bt
}

func (bt *ByteArray) Get_Bytes() []byte {
	return bt.Bytes.Bytes()
}

func (bt *ByteArray) Len() int {
	return bt.Bytes.Len()
}

// Write methods
func (bt *ByteArray) Write_Byte(value byte) *ByteArray {
	bt.Bytes.WriteByte(value)
	return bt
}

func (bt *ByteArray) Write_Short(value uint16) *ByteArray {
	bt.Bytes.Write(binary.BigEndian.AppendUint16([]byte{}, value))
	return bt
}

func (bt *ByteArray) Write_Int(value int) *ByteArray {
	bt.Bytes.Write(binary.BigEndian.AppendUint32([]byte{}, uint32(value)))
	return bt
}

func (bt *ByteArray) Write_Int48(value int64) *ByteArray {
	length := value >> 7

	if length == 0 {
		bt.Bytes.Write([]byte{byte(((value & 127) | 128))})
		bt.Bytes.Write([]byte{byte(0)})
		return bt
	}

	for length != 0 {
		bt.Bytes.Write([]byte{byte(((value & 127) | 128))})
		value = length
		length = length >> 7
	}

	bt.Bytes.Write([]byte{byte(value & 127)})
	return bt
}

func (bt *ByteArray) Write_String(value string) *ByteArray {
	bt.Write_Short(uint16(len(value)))
	bt.Bytes.Write([]byte(value))
	return bt
}

func (bt *ByteArray) Write_Boolean(value bool) *ByteArray {
	if value {
		bt.Bytes.WriteByte(1)
	} else {
		bt.Bytes.WriteByte(0)
	}
	return bt
}

func (bt *ByteArray) Write_Bytes(value []byte) *ByteArray {
	bt.Bytes.Write(value)
	return bt
}

// Read methods
func (bt *ByteArray) Read_Byte() byte {
	value, err := bt.Bytes.ReadByte()

	if err != nil {
		panic(err)
	}

	return value
}

func (bt *ByteArray) Read_Short() uint16 {
	data := make([]byte, 2)
	bt.Bytes.Read(data)
	return binary.BigEndian.Uint16(data)
}

func (bt *ByteArray) Read_Int() int {
	data := make([]byte, 4)
	bt.Bytes.Read(data)
	return int(binary.BigEndian.Uint32(data))
}

func (bt *ByteArray) Read_Int48() int64 {
	var local1 int64 = 0
	var local3 int64 = 0
	var local4 int64 = -1

	for {
		var local2 int64 = int64(bt.Read_Byte())
		local1 = (local1 | ((local2 & 127) << (local3 * 7)))
		local4 = (local4 << 7)
		local3 += 1

		if (local2 & 128) != 128 {
			break
		}
	}

	if ((local4 >> 1) & local1) != 0 {
		local1 = (local1 | (local4))
	}

	return local1
}

func (bt *ByteArray) Read_String() string {
	data := make([]byte, bt.Read_Short())
	bt.Bytes.Read(data)
	return string(data)
}

func (bt *ByteArray) Read_Boolean() bool {
	return bt.Read_Byte() == 1
}

func (bt *ByteArray) Read_Bytes(length int) []byte {
	data := make([]byte, length)
	bt.Bytes.Read(data)
	return data
}
