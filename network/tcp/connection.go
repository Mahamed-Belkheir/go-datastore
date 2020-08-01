package tcp

import (
	"github.com/Mahamed-Belkheir/go-datastore/network"
	"net"
)


type Connection struct {
	conn net.Conn
	SendQueue chan network.Packet
	ReceiveQueue chan network.Packet
	errors chan error
	stopTransmit chan bool
	stopRecieve chan bool
}

func EstablishConnection(conn net.Conn) Connection {
	connection := Connection{
		conn: conn,
		SendQueue: make(chan network.Packet),
		ReceiveQueue: make(chan network.Packet),
		errors: make(chan error),
		stopTransmit: make(chan bool),
		stopRecieve: make(chan bool),
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
		case <- c.stopTransmit:
			return
		}
	}
}

func (c Connection) receieveData() {
	for {
		packet := readTCPPacket(c.conn)
		c.ReceiveQueue <- packet
		select {
		case <- c.stopRecieve:
			return
		}
	}
}

func (c Connection) Stop() {
	c.stopRecieve <- true
	c.stopTransmit <- true
}

func (c Connection) Send(data network.Packet) {
	c.SendQueue <- data
}

func (c Connection) Receive() network.Packet {
	return <- c.ReceiveQueue
}