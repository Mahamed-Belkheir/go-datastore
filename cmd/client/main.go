package main

import (
	c "github.com/Mahamed-Belkheir/go-datastore/network/tcp/client"
	"fmt"
	
	
)

func main() {
	client := c.Client("0.0.0.0:5000", "a", "a")
	conn, err := client.Connect(); if err != nil {
		fmt.Println("error connecting: ", err)
		return
	}
	err = conn.Set("key", "message"); if err != nil {
		fmt.Println("failed to Set, err:", err)
		return
	}
	fmt.Println("done")

}