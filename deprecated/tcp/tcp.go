package tcp

import (
	"fmt"
	t "go-datastore/datastructs"
	"go-datastore/store"
	"io"
	"net"
)

func StartTcpServer(port string, cache *store.Cache) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Could not start TCP server on port", port, "error", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection with error:", err)
		}
		go HandleTcp(conn, cache)
	}
}

func HandleTcp(conn net.Conn, cache *store.Cache) {
	defer conn.Close()
	operation, data, err := t.ParseDataIntoMessage(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(operation, data.Key, data.Data)
	response := cache.Operate(operation, data)
	io.Copy(conn, response.Serialize())
}

/*
Message Structure:
OPERATION - [SET, GET, DELETE] string enum
\n - new line as break point
TYPE - [STRING, COMPLEX] either interpet message as a string or a complex object
\n
DATA - bytes to be stored
\n
*/
