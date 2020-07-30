package datastructs

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"encoding/binary"
)

/*
operation - 1 bytes
datatype - 1 bytes
keysize - 1 byte
key - keysize bytes
datasize - 4 bytes
data - datasize bytes
*/

type Message struct {
	Operation uint8
	DataType  uint8
	Key       string
	Data      []byte
}

func (d *Message) Serialize() *bytes.Buffer {
	var data bytes.Buffer
	data.WriteByte(d.Operation)
	data.WriteByte(d.DataType)

	keySize := uint8(len(d.Key))
	data.WriteByte(keySize)
	data.Write([]byte(d.Key))
	dataSize := uint32(len(d.Data))
	s := make([]byte, 4)
	binary.LittleEndian.PutUint32(s, dataSize)

	data.Write(s)

	data.Write(d.Data)

	return &data
}

func ParseDataIntoMessage(d io.ReadWriter) (string, *Message, error) {
	rw := bufio.NewReadWriter(bufio.NewReader(d), bufio.NewWriter(d))
	
	operation, err := rw.ReadByte(); if err != nil {
		return "", nil, errors.New("Error reading operation")
	}

	dataType, err := rw.ReadByte(); if err != nil {
		return "", nil, errors.New("Error reading Datatype")
	}

	keySize, err := rw.ReadByte(); if err != nil {
		return "", nil, errors.New("Error reading key size")
	}

	keyBytes := make([]byte, keySize)
	if _, err := rw.Read(keyBytes); err != nil {
		return "", nil, errors.New("Error reading key")
	}
	key := string(keyBytes)

	dataSizeBytes := make([]byte, 4)
	if _, err := rw.Read(dataSizeBytes); err != nil {
		return "", nil, errors.New("error reading data size")
	}

	dataSize := binary.LittleEndian.Uint32(dataSizeBytes)

	data := make([]byte, dataSize)
	if _, err := rw.Read(data); err != nil {
		return "", nil, errors.New("error reading data")
	}

	return OperationsMap[operation], &Message{operation, dataType, key, data}, nil
}

func NewMesage(Operation, DataType, key string, data []byte) (*Message, error) {
	op, ok := ReverseOperationsMap[Operation]; if ! ok {
		return nil, errors.New("Invalid Operation " + Operation)
	}
	dataType, ok := ReverseTypesMap[DataType]; if ! ok{
		return nil, errors.New("Invalid DataType "+ DataType)
	}
	return &Message{op, dataType, key, data}, nil
}


var OperationsMap = map[uint8]string {
	0x1: "GET",
	0x2: "SET",
	0x3: "DEL",
}

var ReverseOperationsMap = map[string]uint8 {
	"GET": 0x1,
	"SET": 0x2,
	"DEL": 0x3,
}

var TypesMap = map[uint8]string {
	0x1: "string",
	0x2: "boolean",
	0x3: "integer",
	0x4: "float",
	0x5: "struct",
}

var ReverseTypesMap = map[string]uint8 {
	"string": 0x1,
	"boolean": 0x2,
	"integer": 0x3,
	"float": 0x4,
	"struct": 0x5,
}