package api

import (
	"fmt"
	"go-datastore/serializers"
	"net"
)

func StartTcpServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	fmt.Println("server listening at port", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		b := serializers.TcpData{}
		m := b.Parse(conn)
		fmt.Println(m.Key, string(m.Data))
	}
}

func HandleTcpSet() {}

func HandleTcpGet() {}
