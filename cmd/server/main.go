package main

import (
	"github.com/Mahamed-Belkheir/go-datastore/network/tcp"
)

func main() {
	server := tcp.Server("bob", "password", 5, 10)
	server.Listen("0.0.0.0:5000")
}