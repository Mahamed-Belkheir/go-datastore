package server

import (
	"net"
)

type ConnectionPool struct {
	WorkerPool chan chan net.Conn
	ConnQueue chan net.Conn
	Server *TCPServer
}

func (c ConnectionPool) AddConn(conn net.Conn) {
	c.ConnQueue <- conn
}

func NewConnectionPool(maxWorkers, maxQueue int, server *TCPServer) *ConnectionPool {
	pool := make(chan chan net.Conn, maxWorkers)
	queue := make(chan net.Conn, maxQueue)
	return &ConnectionPool{
		WorkerPool: pool,
		ConnQueue: queue,
		Server: server,
	}
}

func (c ConnectionPool) Initialize() {
	for i := 0; i < cap(c.WorkerPool); i++ {
		worker := NewConnectionWorker(c.Server, c.WorkerPool)
		worker.Start()
	}
	go c.listen()
}

func (c ConnectionPool) listen() {
	for {
		select {
		case conn := <-c.ConnQueue:
			go func(conn net.Conn) {
				connectionPass := <- c.WorkerPool
				connectionPass <- conn
			}(conn)
		}
	}
}


type ConnectionWorker struct {
	Server *TCPServer
	WorkerPool chan chan net.Conn
	WorkerChannel chan net.Conn
	quit chan bool
}

func NewConnectionWorker(server *TCPServer, workerPool chan chan net.Conn) ConnectionWorker {
	return ConnectionWorker{
		Server: server,
		WorkerPool: workerPool,
		WorkerChannel: make(chan net.Conn),
		quit: make(chan bool),
	}
}

func (c ConnectionWorker) Start() {
	// c.Server.e.Publish("log", logging.New("info", "new connection worker started"))
	go func() {
		for {
			c.WorkerPool <- c.WorkerChannel

			select {
			case conn := <- c.WorkerChannel:
				c.Server.handleConn(conn)				
			case <- c.quit:
				return
			}
		}
	}()
}

func (c ConnectionWorker) Stop() {
	go func() {
		c.quit <- true
	}()
}