package tcp

import (
	"net"
)

type Server struct{

}


func (s *Server) Listen() {

}


func (s *Server) handleConn(conn net.Conn) {
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