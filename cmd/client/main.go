package main

import (
	"fmt"
	t "go-datastore/datastructs"
	"io"
	"net"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := t.NewData("string", "user1_token", []byte("USERTOKEN12345")}
	buffer := data.Serialize("GET")
	io.Copy(c, buffer)

	status, newData, err := t.ParseResponse(c)
	if err != nil {
		fmt.Println("error getting data back", err)
	}
	fmt.Println(status, newData)
}
