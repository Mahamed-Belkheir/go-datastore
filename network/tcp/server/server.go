package server

import (
	"fmt"
	"net"
	logging "github.com/Mahamed-Belkheir/go-datastore/logging"
	events "github.com/Mahamed-Belkheir/go-datastore/events"
	t "github.com/Mahamed-Belkheir/go-datastore/network/tcp/utils"
)

type TCPServer struct {
	e *events.EventsBus
	username string
	password string
	Pool *ConnectionPool
}

func Server(username, password string, maxWorkers, maxQueue int) *TCPServer {
	server := &TCPServer{
		e: &events.EventsBus{},
		username: username,
		password: password,
	}
	pool := NewConnectionPool(maxWorkers, maxQueue, server)
	server.Pool = pool
	return server
}

func (s *TCPServer) Listen(address string) {
	s.Pool.Initialize()
	l, err := net.Listen("tcp", address)
	if err != nil {
		s.e.Publish("log", logging.New("error", "tcp server failed to start"))
		panic(err)
	}
	fmt.Println("Server started listening on ", address)
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
	fmt.Println("got a new connection")
	// authenticate
	if  !t.ParseAuth(conn, s.username, s.password) {
		s.e.Publish("log", logging.New("info", fmt.Sprintf("client auth failed. ip: %v", conn.RemoteAddr().String())))
		return
	}

	// loop
	activeConn := t.EstablishConnection(conn)
	
	for {
		select {
		case packet := <- activeConn.ReceiveQueue:
			fmt.Println("Recieved Packet:")
			packet.Operation = "OK"
			activeConn.SendQueue <- packet
		case err := <- activeConn.Errors:
			fmt.Println("Recieved Error:")
			fmt.Println(err)
			return
		}
	}
}
