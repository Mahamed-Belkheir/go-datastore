package main

import (
	"fmt"
	"net"
)

type Message struct {
	Op          uint8
	MessageType uint8
	Data        []byte
}

func (m Message) Serialize() (value []byte) {
	value = append(value, m.Op, m.MessageType)
	value = append(value, m.Data...)
	return value
}

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	m := Message{Op: 1, MessageType: 20, Data: []byte("HELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLOHELLO")}
	c.Write(m.Serialize())
}
