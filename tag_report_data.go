// Copyright (c) 2018 Iori Mizutani
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package llrp

// TagReportData holds an actual parameter in byte and
// how many tags are included in the parameter
type TagReportData struct {
	Data     []byte
	TagCount uint
}

// BuildTagReportDataParameter takes one Tag struct and build TagReportData parameter payload in []byte
// FIXME: Take SetReaderConfig or ROSpec configuration parameters
func NewTagReportDataParam(tag *Tag) []byte {
	// EPCData
	// Calculate the right length fro, epc and pcbits
	epcLengthBits := len(tag.EPC) * 8 // # bytes * 8 = # bits
	length := 4 + 2 + len(tag.EPC)    // header + epcLengthBits + epc
	epcd := EPCData(uint16(length), uint16(epcLengthBits), tag.EPC)

	// ChannlenIndex
	chIndex := ChannelIndex()

	// LastSeenTimeStamp
	timestamp := LastSeenTimestampUTC()

	// TagSeenCount
	tagSeenCount := TagSeenCount()

	// AirProtocolTagData
	aptd := C1G2PC(tag.PCBits)

	//tagReportDataLength := 4 + len(epcd) + len(chIndex) + len(timestamp) + len(tagSeenCount) // Rsvd+Type+length->32bits=4bytes
	tagReportDataLength := len(epcd) + len(aptd) + 4 // Rsvd+Type+length->32bits=4bytes

	// Pack in []byte
	return Pack([]interface{}{
		uint16(240),                 // Rsvd+Type=240 (TagReportData parameter)
		uint16(tagReportDataLength), // Length
		epcd,
		//chIndex,
		//timestamp,
		//tagSeenCount,
		aptd,
	})
}
