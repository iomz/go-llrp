// Copyright (c) 2018 Iori Mizutani
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

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

// Pack the data into (partial) LLRP packet payload.
// TODO: count the data size and return resulting length ?
func Pack(data []interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, v := range data {
		err := binary.Write(buf, binary.BigEndian, v)
		if err != nil {
			panic(err)
		}
	}
	return buf.Bytes()
}

// ReadEvent is the struct to hold data on RFTags
type LLRPReadEvent struct {
	ID []byte
	PC []byte
}

// UnmarshalROAccessReportBody extract ReadEvent from the message value in the ROAccessReport
func UnmarshalROAccessReportBody(roarBody []byte) []*LLRPReadEvent {
	//defer timeTrack(time.Now(), fmt.Sprintf("unpacking %v bytes", len(roarBody)))
	res := []*LLRPReadEvent{}

	// iterate through the parameters in roarBody
	for offset := 0; offset < len(roarBody); {
		parameterType := binary.BigEndian.Uint16(roarBody[offset : offset+2])

		switch parameterType {
		case uint16(240): // TagReportData
			offset += 4
		default:
			offset += int(binary.BigEndian.Uint16(roarBody[offset+2 : offset+4]))
			continue
		}

		// look into TagReportData
		// Now the offset is at the first parameter in the TRD
		var id, pc []byte
		if roarBody[offset] == 141 { // EPC-96
			id = roarBody[offset+1 : offset+13]
			offset += 13
			if roarBody[offset] == 140 { // C1G2-PC parameter
				pc = roarBody[offset+1 : offset+3]
				offset += 3
			}
		} else if binary.BigEndian.Uint16(roarBody[offset:offset+2]) == 241 { // EPCData
			epcDataLength := int(binary.BigEndian.Uint16(roarBody[offset+2 : offset+4])) // length
			//epcLengthBits := binary.BigEndian.Uint16(roarBody[offset+4 : offset+6])      // EPCLengthBits
			id = roarBody[offset+6 : offset+epcDataLength]
			offset += epcDataLength
			if roarBody[offset] == 140 { // C1G2-PC parameter
				pc = roarBody[offset+1 : offset+3]
				offset += 3
			}
		}
		// append the id and pc as an ReadEvent
		res = append(res, &LLRPReadEvent{id, pc})
	}
	return res
}
