package network

type Network interface {
	Client()
	Server()
}

type Server interface {
	Listen()
}

type Client interface {
	Connect()
	Get()
	Set()
	Delete()
}

type Packet struct {
	RequestID uint16
	Operation string
	DataType  string
	Key       string
	Data      []byte
}