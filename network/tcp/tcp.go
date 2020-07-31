package tcp

import (
	events "github.com/Mahamed-Belkheir/go-datastore/events"
)

type TCPNetwork struct{}

func (t *TCPNetwork) Server(username, password string) *TCPServer {
	return &TCPServer{
		e: &events.EventsBus{},
		username: username,
		password: password,
	}
}

func (t *TCPNetwork) Client() *TCPClient {
	return &TCPClient{}
}
