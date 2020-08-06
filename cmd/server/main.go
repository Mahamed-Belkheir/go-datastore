package main

import (
	"github.com/Mahamed-Belkheir/go-datastore/network/tcp/server"
)

func main() {
	server := server.Server("bob", "password", 5, 10)
	server.Listen("0.0.0.0:5000")
}