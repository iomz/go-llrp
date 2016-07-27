package llrp

// Generate Keepalive message.
func Keepalive() []byte {
	var data = []interface{}{
		uint16(HEADER_KA), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),        // Length
		uint32(0),         // ID
	}
	return Pack(data)
}

// Generate KeepaliveAck message.
func KeepaliveAck() []byte {
	var data = []interface{}{
		uint16(HEADER_KAA), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),         // Length
		uint32(0),          // ID
	}
	return Pack(data)
}

// Generate ROAccessReport message.
func ROAccessReport(tagReportData []byte) []byte {
	roAccessReportLength :=
		len(tagReportData) + 10 // Rsvd+Ver+Type+Length+ID->80bits=10bytes
	messageID += 1
	var data = []interface{}{
		uint16(HEADER_ROAR),          // Rsvd+Ver+Type=61 (RO_ACCESS_REPORT)
		uint32(roAccessReportLength), // Message length
		uint32(messageID),            // Message ID
		tagReportData,
	}
	return Pack(data)
}

// Generate ReaderEventNotification message.
func ReaderEventNotification() []byte {
	readerEventNotificationData := ReaderEventNotificationData()
	readerEventNotificationLength :=
		len(readerEventNotificationData) + 10 // Rsvd+Ver+Type+Length+ID->80bits=10bytes
	messageID += 1
	var data = []interface{}{
		uint16(HEADER_REN),                    // Rsvd+Ver+Type=63 (READER_EVENT_NOTIFICATION)
		uint32(readerEventNotificationLength), // Length
		uint32(messageID),                     // ID
		readerEventNotificationData,
	}
	return Pack(data)
}

// Generate SetReaderConfig message.
func SetReaderConfig() []byte {
	keepaliveSpec := KeepaliveSpec()
	setReaderConfigLength :=
		len(keepaliveSpec) + 11 // Rsvd+Ver+Type+Length+ID+R+Rsvd->88bits=11bytes
	messageID += 1
	var data = []interface{}{
		uint16(HEADER_SRC),            // Rsvd+Ver+Type=3 (SET_READER_CONFIG)
		uint32(setReaderConfigLength), // Length
		uint32(messageID),             // ID
		uint8(0),                      // RestoreFactorySetting(no=0)+Rsvd
		keepaliveSpec,
	}
	return Pack(data)
}

// Generate SetReaderConfigResponse message.
func SetReaderConfigResponse() []byte {
	llrpStatus := LLRPStatus()
	setReaderConfigResponseLength :=
		len(llrpStatus) + 10 // Rsvd+Ver+Type+Length+ID+R+Rsvd->80bits=10bytes
	var data = []interface{}{
		uint16(HEADER_SRCR),                   // Rsvd+Ver+Type=13 (SET_READER_CONFIG_RESPONSE)
		uint32(setReaderConfigResponseLength), // Length
		uint32(0), // ID
		llrpStatus,
	}
	return Pack(data)
}
