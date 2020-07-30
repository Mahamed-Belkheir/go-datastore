package tcp

import (
	"net"

	events "github.com/Mahamed-Belkheir/go-datastore/events"
)

type TCPServer struct {
	e *events.EventsBus
}

func (s *TCPServer) Listen() {

}

func (s *TCPServer) handleConn(conn net.Conn) {
	// call log data event
	// authenticate
	// await packets
	// on packet arrival
	// set timeout
	// parse packet
	// call packet event
	// serialize response
	// return response
	// repeat
}
