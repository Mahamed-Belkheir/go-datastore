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
}

type Packet struct {
	Operation string
	DataType  string
	Key       string
	Data      []byte
}