package datastructs

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

type Response struct {
	status  string
	payload *bytes.Buffer
}

func (r *Response) Serialize() *bytes.Buffer {
	var data bytes.Buffer
	data.Write([]byte(r.status))
	data.Write([]byte("\n"))
	io.Copy(&data, r.payload)
	data.Write([]byte("\n"))
	return &data
}

func ParseResponse(conn io.ReadWriter) (string, *Message, error) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	status, err := rw.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return "", nil, errors.New("Error reading message status")
	}
	status = status[:len(status)-1]

	if status != "GET" {
		return status, nil, nil
	}

	_, data, err := ParseDataIntoMessage(rw)

	if err != nil {
		fmt.Println(err)
		return "", nil, errors.New("Error reading message data")
	}
	return status, data, nil
}

func NewResponse(status string, payload *bytes.Buffer) *Response {
	if payload == nil {
		payload = &bytes.Buffer{}
	}
	return &Response{status, payload}
}
