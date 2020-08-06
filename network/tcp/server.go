package tcp

import (
	"fmt"
	"net"
	logging "github.com/Mahamed-Belkheir/go-datastore/logging"
	events "github.com/Mahamed-Belkheir/go-datastore/events"
	// network "github.com/Mahamed-Belkheir/go-datastore/network"
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
	if  !parseAuth(conn, s.username, s.password) {
		s.e.Publish("log", logging.New("info", fmt.Sprintf("client auth failed. ip: %v", conn.RemoteAddr().String())))
		return
	}

	// loop
	activeConn := EstablishConnection(conn)
	
	for {
		select {
		case packet := <- activeConn.ReceiveQueue:
			fmt.Println("Recieved Packet:")
			activeConn.SendQueue <- packet
		case err := <- activeConn.Errors:
			fmt.Println("Recieved Error:")
			fmt.Println(err)
			return
		}
	}
}
