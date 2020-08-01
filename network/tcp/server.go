package tcp

import (
	"fmt"
	"net"
	logging "github.com/Mahamed-Belkheir/go-datastore/logging"
	events "github.com/Mahamed-Belkheir/go-datastore/events"
	network "github.com/Mahamed-Belkheir/go-datastore/network"
)

type TCPServer struct {
	e *events.EventsBus
	username string
	password string
	Pool *ConnectionPool
}

func (s *TCPServer) Listen(address string) {
	s.Pool.Initialize()
	l, err := net.Listen("tcp", address)
	if err != nil {
		s.e.Publish("log", logging.New("error", "tcp server failed to start"))
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			s.e.Publish("log", logging.New("error", "failed to connect to client"))
		}
		s.Pool.AddConn(conn)
	}
}

func (s *TCPServer) handleConn(conn net.Conn) {
	// log new connection
	s.e.Publish("log", logging.New("info", fmt.Sprintf("new client connected. ip: %v", conn.RemoteAddr().String())))
	
	// authenticate
	if  !parseAuth(conn, s.username, s.password) {
		s.e.Publish("log", logging.New("info", fmt.Sprintf("client auth failed. ip: %v", conn.RemoteAddr().String())))
		return
	}

	// loop
	for {
		// get request
		packet := parseTCPPacket(conn)
		// process request
		response := s.handlePacket(packet)
		// send response
		sendData(response, conn)
	}
}

func (s *TCPServer) handlePacket(packet *network.Packet) []byte {
	return []byte{}
}

func sendData(data []byte, conn net.Conn) {
	conn.Write(data)
}