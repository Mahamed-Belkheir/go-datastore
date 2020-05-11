package main

import (
	"fmt"
	"net"
)

type Message struct {
	Op          uint8
	MessageType uint8
	KeySize     uint8
	Key         []byte
	Data        []byte
}

func (m *Message) setData(key string, data []byte) {
	m.Key = []byte(key)
	m.KeySize = uint8(len(m.Key))
	m.Data = data
}

func (m *Message) Serialize() (value []byte) {
	value = append(value, m.Op, m.MessageType, m.KeySize)
	value = append(value, m.Key...)
	value = append(value, m.Data...)
	return value
}

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	m := Message{Op: 1, MessageType: 1}
	m.setData("token", []byte("This is a token"))
	fmt.Println(m.KeySize, m.Key)
	c.Write(m.Serialize())
}
