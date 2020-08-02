package tcp

import (
	"io"
	"bytes"
	"encoding/binary"
	"bufio"
	"net"
	"github.com/Mahamed-Belkheir/go-datastore/network"
)

func readTCPPacket(conn net.Conn) (network.Packet, error) {
	packet, err := UnmarshalTCPPacket(conn); if err != nil {
		return packet, err
	}
	return packet, nil
}

func writeTCPPacket(conn net.Conn, packet network.Packet) error {
	data := MarshalTCPPacket(packet)
	_, err := io.Copy(conn, &data)
	if err != nil {
		return err
	}
	return nil
}


func UnmarshalTCPPacket(conn net.Conn) (packet network.Packet, err error) {
	rw := bufio.NewReader(conn)
	
	id := make([]byte, 2)
	if _, err = rw.Read(id); err != nil {
		return
	}
	packet.RequestID = binary.LittleEndian.Uint16(id)

	op, err := rw.ReadByte(); if err != nil {
		return
	}
	packet.Operation = OperationsMap[op]

	dataType, err := rw.ReadByte(); if err != nil {
		return
	}
	packet.DataType = TypesMap[dataType]

	keySize, err := rw.ReadByte(); if err != nil {
		return 
	}
	keyBytes := make([]byte, keySize)
	if _, err = rw.Read(keyBytes); err != nil {
		return 
	}
	packet.Key = string(keyBytes)

	dataSizeBytes := make([]byte, 4)
	if _, err = rw.Read(dataSizeBytes); err != nil {
		return 
	}
	dataSize := binary.LittleEndian.Uint32(dataSizeBytes)
	data := make([]byte, dataSize)
	if _, err = rw.Read(data); err != nil {
		return 
	}

	packet.Data = data
	return 
}

func MarshalTCPPacket(packet network.Packet) bytes.Buffer {
	var data bytes.Buffer

	id := make([]byte, 2)
	binary.LittleEndian.PutUint16(id, packet.RequestID)
	data.Write(id)

	op := ReverseOperationsMap[packet.Operation]
	data.WriteByte(op)

	dataType := ReverseTypesMap[packet.DataType]
	data.WriteByte(dataType)

	keySize := uint8(len(packet.Key))
	data.WriteByte(keySize)

	key := []uint8(packet.Key)
	data.Write(key)

	dataSize := uint32(len(packet.Data))
	s := make([]byte, 4)
	binary.LittleEndian.PutUint32(s, dataSize)
	data.Write(s)

	data.Write(packet.Data)

	return data
}





func parseAuth(conn net.Conn, username, password string) bool {
	return true
}






var OperationsMap = map[uint8]string {
	0x1: "GET",
	0x2: "SET",
	0x3: "DEL",
}

var ReverseOperationsMap = map[string]uint8 {
	"GET": 0x1,
	"SET": 0x2,
	"DEL": 0x3,
}

var TypesMap = map[uint8]string {
	0x1: "string",
	0x2: "boolean",
	0x3: "integer",
	0x4: "float",
	0x5: "struct",
}

var ReverseTypesMap = map[string]uint8 {
	"string": 0x1,
	"boolean": 0x2,
	"integer": 0x3,
	"float": 0x4,
	"struct": 0x5,
}