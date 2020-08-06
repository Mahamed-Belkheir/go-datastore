package utils

import (
	"fmt"
	"github.com/Mahamed-Belkheir/go-datastore/network"
	"net"
)


type Connection struct {
	conn net.Conn
	SendQueue chan network.Packet
	ReceiveQueue chan network.Packet
	Errors chan error
	stopTransmit chan bool
	stopRecieve chan bool
}

func EstablishConnection(conn net.Conn) Connection {
	fmt.Println("connection established")
	connection := Connection{
		conn: conn,
		SendQueue: make(chan network.Packet),
		ReceiveQueue: make(chan network.Packet),
		Errors: make(chan error),
		stopTransmit: make(chan bool),
		stopRecieve: make(chan bool),
	}
	connection.transmitData()
	connection.receieveData()
	return connection
}

func (c Connection) transmitData() {
	go func() {
		for {
			fmt.Println("looping transmit")
			select {
			case packet := <- c.SendQueue:
				err := writeTCPPacket(c.conn, packet); if err != nil {
					c.Errors <- err
					c.stopRecieve <- true
					return
				}
				fmt.Println("sent data")
	
			case <- c.stopTransmit:
				return
			}
		}
	}()
}

func (c Connection) receieveData() {
	go func(){
		defer fmt.Println("EXITED")
		for {
			fmt.Println("Recieving data")
			packet, err := readTCPPacket(c.conn); if err != nil {
				c.Errors <- err
				c.stopTransmit <- true
				fmt.Println("error reading, exiting")
				return
			}
			fmt.Println("data put in queue")
			c.ReceiveQueue <- packet
			select {
			case <- c.stopRecieve:
				fmt.Println("stopping recieve")
				return
			default:
			}
		}
	}()
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