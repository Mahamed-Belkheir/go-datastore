package types

type TcpMessage struct {
	Op          bool
	MessageType uint8
	Data        []byte
}

func NewTcpMessage(b []byte) *TcpMessage {
	firstByte := uint8(b[0])
	var Op bool
	if firstByte == 1 {
		Op = true
	} else {
		Op = false
	}
	return &TcpMessage{
		Op:          Op,
		MessageType: uint8(b[1]),
		Data:        b[2:],
	}
}
