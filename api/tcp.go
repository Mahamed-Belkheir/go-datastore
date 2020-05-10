package api

import (
	"fmt"
	"net"
)

func StartTcpServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(conn)
	}
}

func HandleTcpSet() {}

func HandleTcpGet() {}
