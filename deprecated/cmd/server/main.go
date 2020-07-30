package main

import (
	"go-datastore/store"
	"go-datastore/tcp"
)

func main() {
	cache := store.NewCache()
	tcp.StartTcpServer(":5000", cache)

}
