package main

import (
	"github.com/Mahamed-Belkheir/go-datastore/network"
	"fmt"
	t "github.com/Mahamed-Belkheir/go-datastore/network/tcp"
	"net"
	
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000"); if err != nil {
		fmt.Println("Could not connect")
		fmt.Println(err)
		return
	}
	activeConn := t.EstablishConnection(conn)
	p := network.Packet{
		RequestID: 1,
		DataType: "string",
		Key: "hello",
		Operation: "SET",
		Data: []uint8("Hello World"),
	}
	p2 := network.Packet{
		RequestID: 2,
		DataType: "string",
		Key: "bye",
		Operation: "SET",
		Data: []uint8("Hello World"),
	}
	activeConn.Send(p)
	activeConn.Send(p2)
	for {}

}