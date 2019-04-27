package llrp

import "time"

// C1G2PC generates C1G2PC parameter from hexpc string.
func C1G2PC(pc uint16) []byte {
	var data = []interface{}{
		uint8(140), // 1+uint7(Type=12)
		pc,         // PC bits
	}
	return Pack(data)
}

// C1G2ReadOpSpecResult generates C1G2ReadOpSpecResult parameter from readData.
func C1G2ReadOpSpecResult(readData []byte) []byte {
	var data = []interface{}{
		uint16(349), // Rsvd+Type=
		uint16(11),  // Length
		uint8(0),    // Result
		uint16(9),   // OpSpecID
		uint16(1),   // ReadDataWordCount
		readData,    // ReadData
	}
	return Pack(data)
}

// ConnectionAttemptEvent generates ConnectionAttemptEvent parameter.
func ConnectionAttemptEvent() []byte {
	var data = []interface{}{
		uint16(256), // Rsvd+Type=256
		uint16(6),   // Length
		uint16(0),   // Status(Success=0)
	}
	return Pack(data)
}

// EPCData generates EPCData parameter from its length and epcLength, and epc.
func EPCData(length uint16, epcLengthBits uint16, epc []byte) []byte {
	var data []interface{}
	if epcLengthBits == 96 {
		data = []interface{}{
			uint8(141), // 1+uint7(Type=13)
			epc,        // 96-bit EPCData string
		}
	} else {
		data = []interface{}{
			uint16(241),           // uint8(0)+uint8(Type=241)
			uint16(length),        // Length
			uint16(epcLengthBits), // EPCLengthBits
			epc, // EPCData string
		}
	}
	return Pack(data)
}

//ChannelIndex : Generates channelIndex for EPC tag
func ChannelIndex() []byte {
	var data = []interface{}{
		uint8(0x87), //type
		uint16(11),  //ch index
	}
	return Pack(data)
}

//LastSeenTimestampUTC : Returns the a UNIX timestamp in microseconds
func LastSeenTimestampUTC() []byte {
	var data = []interface{}{
		uint8(0x84),             //type
		uint64(makeTimestamp()), //timestamp
	}
	return Pack(data)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

//TagSeenCount :
func TagSeenCount() []byte {
	var data = []interface{}{
		uint8(0x88), //type
		uint16(1),   //tag seen count
	}
	return Pack(data)
}

// KeepaliveSpec generates KeepaliveSpec parameter.
func KeepaliveSpec() []byte {
	var data = []interface{}{
		uint16(220),   // Rsvd+Type=220
		uint16(9),     // Length
		uint8(1),      // KeepaliveTriggerType=Periodic(1)
		uint32(10000), // TimeInterval=10000
	}
	return Pack(data)
}

// Status generates LLRPStatus parameter.
func Status() []byte {
	var data = []interface{}{
		uint16(287), // Rsvd+Type=287
		uint16(8),   // Length
		uint16(0),   // StatusCode=M_Success(0)
		uint16(0),   // ErrorDescriptionByteCount=0
	}
	return Pack(data)
}

// PeakRSSI generates PeakRSSI parameter.
func PeakRSSI() []byte {
	var data = []interface{}{
		uint8(134), // 1+uint7(Type=6)
		uint8(203), // PeakRSSI
	}
	return Pack(data)
}

// ReaderEventNotificationData generates ReaderEventNotification parameter.
func ReaderEventNotificationData(currentTime uint64) []byte {
	utcTimeStamp := UTCTimeStamp(currentTime)
	connectionAttemptEvent := ConnectionAttemptEvent()
	readerEventNotificationDataLength := len(utcTimeStamp) +
		len(connectionAttemptEvent) + 4 // Rsvd+Type+length=32bits=4bytes
	var data = []interface{}{
		uint16(246),                               // Rsvd+Type=246 (ReaderEventNotificationData parameter)
		uint16(readerEventNotificationDataLength), // Length
		utcTimeStamp,
		connectionAttemptEvent,
	}
	return Pack(data)
}

/*
// TagReportData generates TagReportData parameter from epcData, peakRSSI, airProtocolTagData, opSpecResult.
func TagReportData(epcData []byte, airProtocolTagData []byte) []byte {
	tagReportDataLength := len(epcData) + len(airProtocolTagData) +
		4 // Rsvd+Type+length->32bits=4bytes
	var data = []interface{}{
		uint16(240),                 // Rsvd+Type=240 (TagReportData parameter)
		uint16(tagReportDataLength), // Length
		epcData,
		airProtocolTagData,
	}
	return Pack(data)
}
*/

// UTCTimeStamp generates UTCTimeStamp parameter at the current time.
func UTCTimeStamp(currentTime uint64) []byte {
	var data = []interface{}{
		uint16(128), // Rsvd+Type=128
		uint16(12),  // Length
		currentTime, // Microseconds
	}
	return Pack(data)
}

//GeneralDeviceCapabilities : Generates General Device Capabilities
func GeneralDeviceCapabilities() []byte {
	numOfAntennas := 52
	//totalReceiveSensitivities := 42
	length := 28 // 28 + totalReceiveSensitivities*8 + 8 + numOfAntennas*9
	var data = []interface{}{
		uint16(137),           //Type 137
		uint16(length),        //Length
		uint16(numOfAntennas), //Max Antenna
		uint16(16384),         //UTC clock support
		uint32(25882),         //Manufacturer
		uint32(2001007),       //Model
		uint32(0x000a352e),
		uint32(0x31342e30),
		uint32(0x2e323430),
	}
	x := Pack(data)

	// for i := 1; i <= totalReceiveSensitivities; i++ {
	// 	x = append(x, ReceiveSensitivityEntry(uint16(i))...)
	// }
	// x = append(x, GPIOCapabilities()...)
	// for i := 1; i <= numOfAntennas; i++ {
	// 	x = append(x, ReceiveSensitivityEntry(uint16(i))...)
	// }

	return x
}

//LlrpCapabilities : generates LLRP_CAPABILITIES
func LlrpCapabilities() []byte {
	var data = []interface{}{
		uint16(142),  //type 142
		uint16(28),   //length
		uint8(72),    //rf survery = no, buffer fille warning = yes, client request opspec = no, tag inventory = no, supoprt event = yes
		uint8(1),     // max priotity level supported
		uint16(0),    //client request opsec timeout
		uint32(1),    // max num of rospec
		uint32(32),   //max num of spec per rospec
		uint32(1),    //max num of inventory spec per AIspec
		uint32(1508), //max num of accessSpec
		uint32(8),    //max num of opspec per AccessSpec
	}
	return Pack(data)
}

//ReguCapabilities : generates Regulatory Capabilities
func ReguCapabilities() []byte {
	var data = []interface{}{
		uint16(143),      //type 143
		uint16(1189 + 8), //length
		uint16(840),      // country code
		uint16(1),        //comm standards, fcc part 15
		UHFCapabilities(52),
	}
	return Pack(data)
}

//C1G2llrpCapabilities : Generates C1G2llrpCapabilities
func C1G2llrpCapabilities() []byte {
	var data = []interface{}{
		uint16(327), //type 327
		uint16(7),   //length
		uint8(64),   //some params
		uint16(2),   //max num of selectec filter per query
	}
	return Pack(data)
}

//GetReaderConfigResponseIdentification : Generate Identification
func GetReaderConfigResponseIdentification() []byte {
	var data = []interface{}{
		uint16(218),        //type
		uint16(17),         //length
		uint8(0),           //id type
		uint16(0),          //byte count
		uint32(0x00080016), //Reader ID
		uint32(0x25ffff11),
		uint16(0xc167),
	}
	return Pack(data)
}

//AntennaProperties :
func AntennaProperties(id uint16) []byte {
	var data = []interface{}{
		uint16(221), //type
		uint16(9),   //length
		uint16(128), //antenna connected
		uint16(id),
		uint16(0), //gain
	}
	return Pack(data)
}

func AntennaConfiguration(id uint16) []byte {
	length := 6 + 6 + 10 + 24 //36
	var data = []interface{}{
		uint16(222), //type
		uint16(length),
		uint16(id),
	}
	x := Pack(data)
	x = append(x, RFReceiver()...)
	x = append(x, RFTransmitter()...)
	x = append(x, C1G2InventoryCommand()...)
	return x
}

func RFReceiver() []byte {
	length := 6
	var data = []interface{}{
		uint16(223), //type
		uint16(length),
		uint16(1), //reciver sensitivity
	}
	return Pack(data)
}

func RFTransmitter() []byte {
	length := 10
	var data = []interface{}{
		uint16(224), //type
		uint16(length),
		uint16(1),  //hop id
		uint16(0),  //ch index
		uint16(81), //tx power
	}
	return Pack(data)
}

func C1G2InventoryCommand() []byte {
	length := 5 + 8 + 11
	var data = []interface{}{
		uint16(330), //type
		uint16(length),
		uint8(0), //Tag inv state aware

	}
	x := Pack(data)
	x = append(x, C1G2RFControl()...)
	x = append(x, C1G2SingulationControl()...)
	return x
}

func C1G2RFControl() []byte {
	length := 8
	var data = []interface{}{
		uint16(335), //type
		uint16(length),
		uint16(1000), //mode index
		uint16(0),    //tari
	}
	return Pack(data)
}

func C1G2SingulationControl() []byte {
	length := 11
	var data = []interface{}{
		uint16(336), //type
		uint16(length),
		uint8(0x40), //session 1
		uint16(32),  //tag pop
		uint32(0),   //tag transit time
	}
	return Pack(data)
}

// UHFCapabilities :
func UHFCapabilities(numOfAntennas int) []byte {
	numOfPowerLevel := 81
	length := 1189 //numOfPowerLevel*8 + 213 + 324 + 4 //transmitpowerlevel + freqinfo + c1g2 + type + length
	var data = []interface{}{
		uint16(144),    //type
		uint16(length), //length
	}
	x := Pack(data)

	for i := 1; i <= numOfPowerLevel; i++ {
		x = append(x, TransmitPowerLevelEntry(uint16(i), uint16(1000+25*(i-1)))...)
	}
	x = append(x, FrequencyInformation()...)
	x = append(x, C1G2UHFModeRFTable()...)
	return x
}

//TransmitPowerLevelEntry :
func TransmitPowerLevelEntry(id uint16, powerLevel uint16) []byte {
	var data = []interface{}{
		uint16(145),        //type
		uint16(8),          //length
		uint16(id),         //id
		uint16(powerLevel), //power value
	}
	return Pack(data)
}

//FrequencyInformation :
func FrequencyInformation() []byte {
	length := 213 //2 + 2 + 1 + 208
	var data = []interface{}{
		uint16(146),    //type
		uint16(length), //length
		uint8(1),       //hopping
	}
	x := Pack(data)
	x = append(x, FrequencyHopTable()...)
	return x
}

//FrequencyHopTable :
func FrequencyHopTable() []byte {
	numOfHops := 50
	length := 208 //2 + 2 + 1 + 1 + 2 + numOfHops*4
	var data = []interface{}{
		uint16(147),       //type
		uint16(length),    //length
		uint8(1),          // hop table id
		uint8(0),          // reserved
		uint16(numOfHops), //num of hops
	}
	x := Pack(data)
	for i := 0; i < numOfHops; i++ {
		x = append(x, frequency(903250+i)...) //some random frequencies\
	}
	return x
}

func frequency(frequency int) []byte {
	var data = []interface{}{
		uint32(frequency),
	}
	return Pack(data)
}

//C1G2UHFModeRFTable :
func C1G2UHFModeRFTable() []byte {

	length := 324 //2 + 2 + 32*10
	var data = []interface{}{
		uint16(328),    //type
		uint16(length), //length
	}
	x := Pack(data)
	x = append(x, C1G2UHFModeRFTableEntry(0)...)
	x = append(x, C1G2UHFModeRFTableEntry(1)...)
	x = append(x, C1G2UHFModeRFTableEntry(2)...)
	x = append(x, C1G2UHFModeRFTableEntry(3)...)
	x = append(x, C1G2UHFModeRFTableEntry(4)...)
	x = append(x, C1G2UHFModeRFTableEntry(1000)...)
	x = append(x, C1G2UHFModeRFTableEntry(1001)...)
	x = append(x, C1G2UHFModeRFTableEntry(1002)...)
	x = append(x, C1G2UHFModeRFTableEntry(1003)...)
	x = append(x, C1G2UHFModeRFTableEntry(1004)...)
	return x
}

//C1G2UHFModeRFTableEntry :
func C1G2UHFModeRFTableEntry(mode int) []byte {
	var data = []interface{}{
		uint16(329),    //type
		uint16(32),     //length
		uint32(mode),   //mode identifier
		uint8(80),      //dr : yes , EPC HAG T&C confromance : no
		uint8(0),       //m : 0
		uint8(2),       // forward link mod
		uint8(2),       //speectral mask indicator
		uint32(640000), // bdr
		uint32(1500),   //PIE
		uint32(6250),   //max tari
		uint32(6250),   //min tari
		uint32(0),      //tari step
	}
	return Pack(data)
}
