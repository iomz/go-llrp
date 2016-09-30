package llrp

import (
	"bytes"
	"encoding/binary"
)

// LLRP header values
const (
	ROAccessReportHeader          = 1085
	ReaderEventNotificationHeader = 1087
	SetReaderConfigHeader         = 1027
	SetReaderConfigResponseHeader = 1037
	KeepaliveHeader               = 1086
	KeepaliveAckHeader            = 1096
)

// Check if error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Pack the data into (partial) LLRP packet payload.
// TODO: count the data size and return resulting length
func Pack(data []interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, v := range data {
		err := binary.Write(buf, binary.BigEndian, v)
		check(err)
	}
	return buf.Bytes()
}
