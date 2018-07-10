package llrp

// Keepalive generates Keepalive message.
func Keepalive() []byte {
	var data = []interface{}{
		uint16(KeepaliveHeader), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),              // Length
		uint32(0),               // ID
	}
	return Pack(data)
}

// KeepaliveAck generates KeepaliveAck message.
func KeepaliveAck() []byte {
	var data = []interface{}{
		uint16(KeepaliveAckHeader), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),                 // Length
		uint32(0),                  // ID
	}
	return Pack(data)
}

// ReaderEventNotification generates ReaderEventNotification message.
func ReaderEventNotification(messageID uint32, currentTime uint64) []byte {
	readerEventNotificationData := ReaderEventNotificationData(currentTime)
	readerEventNotificationLength :=
		len(readerEventNotificationData) + 10 // Rsvd+Ver+Type+Length+ID->80bits=10bytes
	var data = []interface{}{
		uint16(ReaderEventNotificationHeader), // Rsvd+Ver+Type=63 (READER_EVENT_NOTIFICATION)
		uint32(readerEventNotificationLength), // Length
		messageID, // ID
		readerEventNotificationData,
	}
	return Pack(data)
}

// SetReaderConfig generates SetReaderConfig message.
func SetReaderConfig(messageID uint32) []byte {
	keepaliveSpec := KeepaliveSpec()
	setReaderConfigLength :=
		len(keepaliveSpec) + 11 // Rsvd+Ver+Type+Length+ID+R+Rsvd->88bits=11bytes
	var data = []interface{}{
		uint16(SetReaderConfigHeader), // Rsvd+Ver+Type=3 (SET_READER_CONFIG)
		uint32(setReaderConfigLength), // Length
		messageID,                     // ID
		uint8(0),                      // RestoreFactorySetting(no=0)+Rsvd
		keepaliveSpec,
	}
	return Pack(data)
}

// SetReaderConfigResponse generates SetReaderConfigResponse message.
func SetReaderConfigResponse() []byte {
	llrpStatus := Status()
	setReaderConfigResponseLength :=
		len(llrpStatus) + 10 // Rsvd+Ver+Type+Length+ID+R+Rsvd->80bits=10bytes
	var data = []interface{}{
		uint16(SetReaderConfigResponseHeader), // Rsvd+Ver+Type=13 (SET_READER_CONFIG_RESPONSE)
		uint32(setReaderConfigResponseLength), // Length
		uint32(0), // ID
		llrpStatus,
	}
	return Pack(data)
}
