package llrp

// Keepalive returns the encoded KEEPALIVE message with the provided message ID.
func Keepalive(messageID uint32) []byte {
	var data = []interface{}{
		uint16(KeepaliveHeader), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),              // Length
		messageID,               // ID
	}
	return Pack(data)
}

// KeepaliveAck returns the encoded KEEPALIVE_ACK message with the provided message ID.
func KeepaliveAck(messageID uint32) []byte {
	var data = []interface{}{
		uint16(KeepaliveAckHeader), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),                 // Length
		messageID,                  // ID
	}
	return Pack(data)
}

// ReaderEventNotification builds a READER_EVENT_NOTIFICATION message for the given timestamp.
func ReaderEventNotification(messageID uint32, currentTime uint64) []byte {
	readerEventNotificationData := ReaderEventNotificationData(currentTime)
	readerEventNotificationLength :=
		len(readerEventNotificationData) + 10 // Rsvd+Ver+Type+Length+ID->80bits=10bytes
	var data = []interface{}{
		uint16(ReaderEventNotificationHeader), // Rsvd+Ver+Type=63 (READER_EVENT_NOTIFICATION)
		uint32(readerEventNotificationLength), // Length
		messageID,                             // ID
		readerEventNotificationData,
	}
	return Pack(data)
}

// SetReaderConfig builds a SET_READER_CONFIG message that enables periodic keepalives.
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

// SetReaderConfigResponse returns a SET_READER_CONFIG_RESPONSE message with a success status.
func SetReaderConfigResponse(messageID uint32) []byte {
	llrpStatus := Status()
	setReaderConfigResponseLength :=
		len(llrpStatus) + 10 // Rsvd+Ver+Type+Length+ID+R+Rsvd->80bits=10bytes
	var data = []interface{}{
		uint16(SetReaderConfigResponseHeader), // Rsvd+Ver+Type=13 (SET_READER_CONFIG_RESPONSE)
		uint32(setReaderConfigResponseLength), // Length
		messageID,                             // ID
		llrpStatus,
	}
	return Pack(data)
}

// GetReaderCapability requests all reader capabilities from the device.
func GetReaderCapability(messageID uint32) []byte {
	getReaderCapabilityLength := 1 + 10
	var data = []interface{}{
		uint16(GetReaderCapabilityHeader),
		uint32(getReaderCapabilityLength),
		messageID,
		uint8(0), //all capabilities
	}
	return Pack(data)
}

// GetReaderCapabilityResponse returns a GET_READER_CAPABILITIES_RESPONSE populated with static data.
func GetReaderCapabilityResponse(messageID uint32) []byte {

	llrpStatus := Status()
	generalCapabilites := GeneralDeviceCapabilities()
	llrpCapabilities := LlrpCapabilities()
	c1g2llrpCapabilities := C1G2llrpCapabilities()
	reguCapabilitles := ReguCapabilities()
	length := 2 + 4 + 4 + len(llrpStatus) + len(generalCapabilites) + len(llrpCapabilities) + len(reguCapabilitles) + len(c1g2llrpCapabilities)
	var data = []interface{}{
		uint16(GetReaderCapabilityResponseHeader),
		uint32(length),
		uint32(messageID),
		llrpStatus,
		generalCapabilites,
		llrpCapabilities,
		reguCapabilitles,
		// uint8(0),
		// uint8(0),
		// uint8(0),
		c1g2llrpCapabilities,
	}
	return Pack(data)
}

// GetReaderConfigResponse returns a GET_READER_CONFIG_RESPONSE containing basic reader identity information.
func GetReaderConfigResponse(messageID uint32) []byte {
	llrpStatus := Status()
	//numOfAntennas := 52
	identification := GetReaderConfigResponseIdentification()
	length := 2 + 4 + 4 + len(llrpStatus) + len(identification) //+ numOfAntennas*9 + numOfAntennas*36
	var data = []interface{}{
		uint16(GetReaderConfigResponseHeader),
		uint32(length),
		messageID,
		llrpStatus,
		identification,
	}
	// x := Pack(data)
	// for i := 1; i <= numOfAntennas; i++ {
	// 	x = append(x, AntennaProperties(uint16(i))...)
	// }
	// for i := 1; i <= numOfAntennas; i++ {
	// 	x = append(x, AntennaConfiguration(uint16(i))...)
	// }
	return Pack(data)
}

// DeleteAccessSpecResponse returns a DELETE_ACCESS_SPEC_RESPONSE with a success status.
func DeleteAccessSpecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint16(DeleteAccessSpecResponseHeader),
		uint32(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

// DeleteRospecResponse returns a DELETE_ROSPEC_RESPONSE with a success status.
func DeleteRospecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint16(DeleteRospecResponseHeader),
		uint32(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

// AddRospecResponse returns an ADD_ROSPEC_RESPONSE with a success status.
func AddRospecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint16(AddRospecResponseHeader),
		uint32(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

// EnableRospecResponse returns an ENABLE_ROSPEC_RESPONSE with a success status.
func EnableRospecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint16(EnableRospecResponseHeader),
		uint32(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

// ReceiveSensitivityEntries builds the receive sensitivity list used in general capabilities.
func ReceiveSensitivityEntries(numOfAntennas int) []interface{} {
	var data = []interface{}{}
	for i := 1; i <= numOfAntennas; i++ {
		x := ReceiveSensitivityEntry(uint16(i))
		data = append(data, x)
	}
	return data
}

// ReceiveSensitivityEntry returns a single receive sensitivity entry for the given antenna ID.
func ReceiveSensitivityEntry(id uint16) []byte {
	var data = []interface{}{
		uint16(139), //type
		uint16(8),   //length
		uint16(id),  //id
		uint16(11),  //receive sentitvitiy value
	}
	return Pack(data)
}

// GPIOCapabilities returns the GPIO capabilities parameter declaring zero ports.
func GPIOCapabilities() []byte {
	var data = []interface{}{
		uint16(141), //type
		uint16(8),   //length
		uint16(0),   //num of GPI port
		uint16(0),   //num of GPO port
	}
	return Pack(data)
}

// AntennaAirPortList returns the antenna air protocol list for the supplied antenna count.
func AntennaAirPortList(numOfAntennas int) []interface{} {
	var data = []interface{}{}
	for i := 1; i <= numOfAntennas; i++ {
		x := AntennaAirPort(uint16(i))
		data = append(data, x)
	}
	return data
}

// AntennaAirPort returns a single antenna air protocol entry for the given antenna ID.
func AntennaAirPort(id uint16) []byte {
	var data = []interface{}{
		uint16(140), //type
		uint16(9),   //length
		uint16(id),
		uint16(1), //num of protocols
		uint8(1),  //protocol id : EPCGlobal Class 1 Gen 2
	}
	return Pack(data)
}

// ImpinjEnableCutomMessage returns the vendor-specific IMPINJ enable message.
func ImpinjEnableCutomMessage(id uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint16(ImpinjEnableCutomMessageHeader), //type
		uint32(23),                             //length
		uint32(id),                             //id
		uint32(25822),                          //vendor id
		uint8(22),                              //subtype
		llrpStatus,
	}
	return Pack(data)
}
