package llrp

import (
	"bytes"
	"encoding/binary"
)

const (
	HEADER_ROAR = 1085
	HEADER_REN  = 1087
	HEADER_SRC  = 1027
	HEADER_SRCR = 1037
	HEADER_KA   = 1086
	HEADER_KAA  = 1096
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
