package tcp

import (
	events "github.com/Mahamed-Belkheir/go-datastore/events"
)

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

func Client() *TCPClient {
	return &TCPClient{}
}
