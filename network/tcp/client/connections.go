package client

import (
	"github.com/Mahamed-Belkheir/go-datastore/network/tcp/utils"
	"github.com/Mahamed-Belkheir/go-datastore/network"
)


type ConnectionManager struct {
	Conn *utils.Connection
	Messages map[uint16] chan network.Packet
}

func NewClientConnectionManager(conn *utils.Connection) ConnectionManager {
	return ConnectionManager{
		Conn: conn,
	}
}

func (c *ConnectionManager) RecievedPacket(id uint16) {
	delete(c.Messages, id)
}

func (c *ConnectionManager) SendPacket(packet network.Packet, reciever chan network.Packet) {
	id := packet.RequestID
	c.Messages[id] = reciever
	c.Conn.SendQueue <- packet
}

func (c *ConnectionManager) listen() {
	for {
		packet := <- c.Conn.ReceiveQueue
		receiver := c.Messages[packet.RequestID]
		receiver <- packet
	}
}