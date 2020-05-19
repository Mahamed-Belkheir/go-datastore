package datastructs

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"unsafe"
)

/*
operation - 1 bytes
datatype - 1 bytes
keysize - 1 byte
key - keysize bytes
datasize - 4 bytes
data - datasize bytes
*/

type Data struct {
	Operation uint8
	DataType  uint8
	Key       []byte
	Data      []byte
}

func (d *Data) Serialize(operation string) *bytes.Buffer {
	var data bytes.Buffer
	data.WriteByte(d.Operation)
	data.WriteByte(d.DataType)
	keySize := uint8(len(d.Key))
	data.WriteByte(keySize)
	dataSize := uint32(len(d.Data))
	data.Write((*[4]byte)(unsafe.Pointer(&dataSize))[:])
	data.Write([]byte(d.Key))
	data.Write(d.Data)
	return &data
}

func ParseTCPMessage(conn io.ReadWriter) (string, *Data, error) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	operation, err := rw.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return "", nil, errors.New("Error reading message operation")
	}
	operation = operation[:len(operation)-1]

	dataType, err := rw.ReadString('\n')
	if err != nil {
		return "", nil, errors.New("Error reading message type")
	}
	dataType = dataType[:len(dataType)-1]

	key, err := rw.ReadString('\n')
	if err != nil {
		return "", nil, errors.New("Error reading message key")
	}
	key = key[:len(key)-1]

	data, err := rw.ReadBytes(byte('\n'))
	if err != nil {
		return "", nil, errors.New("Error reading message data")
	}
	data = data[:len(data)-1]

	return operation, NewData(dataType, key, data), nil
}

func NewData(DataType, key string, data []byte) *Data {
	return &Data{DataType, key, data}
}
