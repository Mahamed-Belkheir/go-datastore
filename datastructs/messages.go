package datastructs

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

type Data struct {
	DataType string
	Key      string
	Data     []byte
}

func (d *Data) Serialize(operation string) *bytes.Buffer {
	var data bytes.Buffer
	data.Write([]byte(operation + "\n"))
	data.Write([]byte(d.DataType + "\n"))
	data.Write([]byte(d.Key + "\n"))
	data.Write(d.Data)
	data.Write([]byte("\n"))
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

	return operation, ParseData(dataType, key, data), nil
}

func ParseData(DataType, key string, data []byte) *Data {
	return &Data{DataType, key, data}
}
