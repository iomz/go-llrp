package llrp

// Generate Keepalive message.
func Keepalive() []byte {
	var data = []interface{}{
		uint16(H_Keepalive), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),        // Length
		uint32(0),         // ID
	}
	return pack(data)
}

// Generate KeepaliveAck message.
func KeepaliveAck() []byte {
	var data = []interface{}{
		uint16(H_KeepaliveAck), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),         // Length
		uint32(0),          // ID
	}
	return pack(data)
}

// Generate ROAccessReport message.
func ROAccessReport(tagReportData []byte, messageID int) []byte {
	roAccessReportLength :=
		len(tagReportData) + 10 // Rsvd+Ver+Type+Length+ID->80bits=10bytes
	var data = []interface{}{
		uint16(H_ROAccessReport),          // Rsvd+Ver+Type=61 (RO_ACCESS_REPORT)
		uint32(roAccessReportLength), // Message length
		uint32(messageID),            // Message ID
		tagReportData,
	}
	return pack(data)
}

// Generate ReaderEventNotification message.
func ReaderEventNotification(messageID int) []byte {
	readerEventNotificationData := ReaderEventNotificationData()
	readerEventNotificationLength :=
		len(readerEventNotificationData) + 10 // Rsvd+Ver+Type+Length+ID->80bits=10bytes
	var data = []interface{}{
		uint16(H_ReaderEventNotification),                    // Rsvd+Ver+Type=63 (READER_EVENT_NOTIFICATION)
		uint32(readerEventNotificationLength), // Length
		uint32(messageID),                     // ID
		readerEventNotificationData,
	}
	return pack(data)
}

// Generate SetReaderConfig message.
func SetReaderConfig(messageID int) []byte {
	keepaliveSpec := KeepaliveSpec()
	setReaderConfigLength :=
		len(keepaliveSpec) + 11 // Rsvd+Ver+Type+Length+ID+R+Rsvd->88bits=11bytes
	var data = []interface{}{
		uint16(H_SetReaderConfig),            // Rsvd+Ver+Type=3 (SET_READER_CONFIG)
		uint32(setReaderConfigLength), // Length
		uint32(messageID),             // ID
		uint8(0),                      // RestoreFactorySetting(no=0)+Rsvd
		keepaliveSpec,
	}
	return pack(data)
}

// Generate SetReaderConfigResponse message.
func SetReaderConfigResponse() []byte {
	llrpStatus := LLRPStatus()
	setReaderConfigResponseLength :=
		len(llrpStatus) + 10 // Rsvd+Ver+Type+Length+ID+R+Rsvd->80bits=10bytes
	var data = []interface{}{
		uint16(H_SetReaderConfigResponse),                   // Rsvd+Ver+Type=13 (SET_READER_CONFIG_RESPONSE)
		uint32(setReaderConfigResponseLength), // Length
		uint32(0), // ID
		llrpStatus,
	}
	return pack(data)
}
