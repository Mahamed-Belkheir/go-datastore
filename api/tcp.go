package api

import (
	"fmt"
	"go-datastore/serializers"
	"go-datastore/store"
	"net"
)

func StartTcpServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	fmt.Println("server listening at port", port)
	cache := store.NewCache()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}

		b := serializers.TcpData{}
		fmt.Println("before parse")
		m := b.Parse(conn)
		fmt.Println("we're outside")

		if m.Op == uint8(1) {
			fmt.Println("we're inside")
			result := cache.Get(m.Key)
			conn.Write(result.Data)
			fmt.Println(result.Data)
			conn.Close()
		} else {
			fmt.Println("we're in else")
			cache.Set(m.Key, m)
		}

	}
}

func HandleTcpSet() {}

func HandleTcpGet() {}
