package tcp

import (
	"github.com/Mahamed-Belkheir/go-datastore/network"
	"net"
)


type Connection struct {
	conn net.Conn
	SendQueue chan network.Packet
	ReceiveQueue chan network.Packet
}

func EstablishConnection(conn net.Conn) Connection {
	connection := Connection{
		conn: conn,
		SendQueue: make(chan network.Packet),
	}
	go connection.transmitData()
	go connection.receieveData()
	return connection
}

func (c Connection) transmitData() {
	for {
		select {
		case packet := <- c.SendQueue:
			writeTCPPacket(c.conn, packet)
		}
	}
}

func (c Connection) receieveData() {
	for {
		packet := readTCPPacket(c.conn)
		c.ReceiveQueue <- packet
	}
}

func (c Connection) Send(data network.Packet) {
	c.SendQueue <- data
}

func (c Connection) Receive() network.Packet {
	return <- c.ReceiveQueue
}