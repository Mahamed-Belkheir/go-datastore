package tcp

import (
	"fmt"
	"net"
	logging "github.com/Mahamed-Belkheir/go-datastore/logging"
	events "github.com/Mahamed-Belkheir/go-datastore/events"
)

type TCPServer struct {
	e *events.EventsBus
	username string
	password string
}

func (s *TCPServer) Listen(address string) {
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
		go s.handleConn(conn)
	}
}

func (s *TCPServer) handleConn(conn net.Conn) {
	s.e.Publish("log", logging.New("info", fmt.Sprintf("new client connected. ip: %v", conn.RemoteAddr().String())))
	if  !parseAuth(conn, s.username, s.password) {
		s.e.Publish("log", logging.New("info", fmt.Sprintf("client auth failed. ip: %v", conn.RemoteAddr().String())))
		return
	}
	//	loop
		// 	get request
		// 	send response
}
