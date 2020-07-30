package tcp

import (
	events "github.com/Mahamed-Belkheir/go-datastore/events"
)

type TCPNetwork struct{}

func (t *TCPNetwork) Server() *TCPServer {
	return &TCPServer{
		e: &events.EventsBus{},
	}
}

func (t *TCPNetwork) Client() *TCPClient {
	return &TCPClient{}
}
