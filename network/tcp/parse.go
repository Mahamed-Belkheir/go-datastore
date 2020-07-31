package tcp

import (
	"net"
	"github.com/Mahamed-Belkheir/go-datastore/network"
)

func parseTCPPacket() *network.Packet {
	var packet network.Packet

	return &packet
}

func parseAuth(conn net.Conn, username, password string) bool {
	return true
}