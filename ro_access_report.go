// Copyright (c) 2018 Iori Mizutani
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package llrp

import (
	"net"
)

// ROAccessReport is a struct for ROAR message
type ROAccessReport struct {
	length int
	data   []byte
}

// Send the ROAR throgh the conn
func (roar *ROAccessReport) Send(conn net.Conn) error {
	// Send
	_, err := conn.Write(roar.data)
	if err != nil {
		return err
	}
	return nil
}

// NewROAccessReport returns a pointer to a new ROAccessReport message.
func NewROAccessReport(tagReportData []byte, messageID uint32) *ROAccessReport {
	roar := &ROAccessReport{}
	roar.length = len(tagReportData) + 10 // Rsvd+Ver+Type+Length+ID->80bits=10bytes
	roar.data = Pack([]interface{}{
		uint16(ROAccessReportHeader), // Rsvd+Ver+Type=61 (RO_ACCESS_REPORT)
		uint32(roar.length),          // Message length
		messageID,                    // Message ID
		tagReportData,
	})
	return roar
}
