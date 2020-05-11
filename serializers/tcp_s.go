package serializers

import (
	"fmt"
	"go-datastore/types"
	"net"
)

type TcpData []byte

func (t TcpData) Parse(c net.Conn) *types.TcpMessage {
	count := 8
	for {
		b := make([]byte, count*count)
		n, err := c.Read(b)
		if err != nil {
			fmt.Println("failed to read message", err)
			break
		}
		t = append(t, b[:n]...)
		fmt.Println(count*count, count, n)
		count = count * 2
	}
	return types.NewTcpMessage(t)
}
