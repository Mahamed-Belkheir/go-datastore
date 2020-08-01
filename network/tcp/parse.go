package tcp

import (
	"net"
	"github.com/Mahamed-Belkheir/go-datastore/network"
)

func readTCPPacket(conn net.Conn) network.Packet {
	var packet network.Packet

	return packet
}

func writeTCPPacket(conn net.Conn, packet network.Packet) {
	conn.Write([]byte{})
}

func parseAuth(conn net.Conn, username, password string) bool {
	return true
}