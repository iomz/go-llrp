package llrp

import (
	"bytes"
	"encoding/binary"
)

const (
	H_ROAccessReport = 1085
	H_ReaderEventNotification  = 1087
	H_SetReaderConfig  = 1027
	H_SetReaderConfigResponse = 1037
	H_Keepalive   = 1086
	H_KeepaliveAck  = 1096
)

// Check if error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Pack the data into (partial) LLRP packet payload.
// TODO: count the data size and return resulting length
func pack(data []interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, v := range data {
		err := binary.Write(buf, binary.BigEndian, v)
		check(err)
	}
	return buf.Bytes()
}
