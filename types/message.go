package types

type TcpMessage struct {
	Op          uint8
	MessageType uint8
	Key         string
	Data        []byte
}

func NewTcpMessage(b []byte) *TcpMessage {

	return &TcpMessage{
		Op:          b[0],
		MessageType: uint8(b[1]),
		Key:         string(b[3 : 3+int(b[2])]),
		Data:        b[3+int(b[2]):],
	}
}
