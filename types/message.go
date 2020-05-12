package types

type TcpMessage struct {
	Op       uint8
	DataType uint8
	Key      string
	Data     []byte
}

func NewTcpMessage(b []byte) *TcpMessage {

	return &TcpMessage{
		Op:       b[0],
		DataType: uint8(b[1]),
		Key:      string(b[3 : 3+int(b[2])]),
		Data:     b[3+int(b[2]):],
	}
}

var DataTypes = map[uint8]string{
	1: "text",
	2: "boolean",
	3: "integer",
	4: "blob",
	5: "map",
}
