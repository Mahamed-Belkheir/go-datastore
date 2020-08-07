package client

import (
	"errors"
	"github.com/Mahamed-Belkheir/go-datastore/network/tcp/utils"
	"github.com/Mahamed-Belkheir/go-datastore/network"
	"net"
)

type TCPClient struct {
	host string
	username string
	password string
}

type TCPClientInstance struct {
	Conn ConnectionManager
}

func Client(host, username, password string) *TCPClient {
	return &TCPClient{
		host: host,
		username: username,
		password: password,
	}
}

func (c TCPClient) Connect() (*TCPClientInstance, error) {
	conn, err := net.Dial("tcp", c.host); if err != nil {
		return nil, err
	}
	activeConn := utils.EstablishConnection(conn)
	return &TCPClientInstance{
		Conn: NewClientConnectionManager(&activeConn),
	}, nil
}


func (c TCPClientInstance) Get(key string) (value interface{}, err error) {

	return
}

func (c TCPClientInstance) Set(key string, value interface{}) (err error) {
	packet, err := composePacket(key, value); if err != nil {
		return err
	}
	var resultChan chan network.Packet
	c.Conn.SendPacket(*packet, resultChan)
	result := <- resultChan
	if result.Operation != "OK" {
		return errors.New(string(result.Data))
	}
	return nil
}

func composePacket(key string, value interface{}) (*network.Packet, error) {
	// get unique request id
	requestId := getRequestId()
	data, dataType, err := utils.Serialize(value); if err != nil {
		return nil, err
	}
	return &network.Packet{
		RequestID: requestId,
		Operation: "SET",
		Key: key,
		DataType: dataType,
		Data: data,
	}, nil
} 



func getRequestId() uint16 {
	return 1
}