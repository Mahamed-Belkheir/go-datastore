package utils

import (
	"encoding/json"
	"math"
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





func ParseAuth(conn net.Conn, username, password string) bool {
	return true
}






var OperationsMap = map[uint8]string {
	0x1: "OK",
	0x2: "PING",
	0x3: "GET",
	0x4: "SET",
	0x5: "DEL",
	0x6: "ERR",
}

var ReverseOperationsMap = map[string]uint8 {
	"OK": 	0x1,
	"PING": 0x2,
	"GET": 	0x3,
	"SET": 	0x4,
	"DEL": 	0x5,
	"ERR":  0x6,
}

var TypesMap = map[uint8]string {
	0x1: "string",
	0x2: "boolean",
	0x3: "integer",
	0x4: "float",
	0x5: "json",
}

var ReverseTypesMap = map[string]uint8 {
	"string": 	0x1,
	"boolean": 	0x2,
	"integer": 	0x3,
	"float": 	0x4,
	"json": 	0x5,
}


func Serialize(data interface{}) ([]byte, string, error) {
	switch data.(type) {
	case string:
		return []byte(data.(string)), "string", nil
	case int:
		result := make([]byte, 8)
		binary.LittleEndian.PutUint64(result, uint64(data.(int)))
		return result, "integer", nil
	case bool:
		if data.(bool) {
			return []byte{0x1}, "boolean", nil
		} else {
			return []byte{0x0}, "boolean", nil
		}
	case float64:
		uintrep := math.Float64bits(data.(float64))
		result := make([]byte, 8)
		binary.LittleEndian.PutUint64(result, uint64(uintrep))
		return result, "float", nil
	default:
		result, err:= json.Marshal(data); if err != nil {
			return nil, "err", err
		}
		return result, "json", nil
	}
}

func Deserialize(data []byte, dataType string) (interface{}, error) {
	switch dataType {
	case "string":
		return DeserializeString(data), nil
	case "integer":
		return DeserializeInt(data), nil
	case "boolean":
		return DeserializeBoolean(data), nil
	case "float":
		return DeserializeFloat(data), nil
	default:
		var result interface{} 
		err := DeserializeJson(data, &result)
		return result, err
	}
}

func DeserializeString(data []byte) string {
	return string(data)
}

func DeserializeInt(data []byte) int64 {
	return int64(binary.LittleEndian.Uint64(data))
}

func DeserializeBoolean(data []byte) bool {
	if data[0] == 1 {
		return true
	} else {
		return false
	}
}

func DeserializeFloat(data []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(data))
}

func DeserializeJson(data []byte, target *interface{}) error {
	err := json.Unmarshal(data, target); if err != nil {
		return err
	}
	return nil
}