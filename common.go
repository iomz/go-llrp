package llrp

import (
	"bytes"
	"encoding/binary"
)

// Check if error
func Check(e error) {
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
		Check(err)
	}
	return buf.Bytes()
}
