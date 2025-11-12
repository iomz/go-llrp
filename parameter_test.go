package llrp

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"testing"
)

func TestC1G2PC(t *testing.T) {
	var b, out []byte
	b = C1G2PC(12288)
	out = []byte{140, 48, 0}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestC1G2ReadOpSpecResult(t *testing.T) {
	var b, out, dummy []byte
	b = C1G2ReadOpSpecResult(dummy)
	out = []byte{1, 93, 0, 11, 0, 0, 9, 0, 1}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestConnectionAttemptEvent(t *testing.T) {
	var b, out []byte
	b = ConnectionAttemptEvent()
	out = []byte{1, 0, 0, 6, 0, 0}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestEPCData(t *testing.T) {
	var b, out []byte
	epc, _ := hex.DecodeString("302DB319A000004000000003")
	b = EPCData(18, 96, epc)
	out = []byte{141, 48, 45, 179, 25, 160, 0, 0, 64, 0, 0, 0, 3}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestKeepaliveSpec(t *testing.T) {
	var b, out []byte
	b = KeepaliveSpec()
	out = []byte{0, 220, 0, 9, 1, 0, 0, 39, 16}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestStatus(t *testing.T) {
	var b, out []byte
	b = Status()
	out = []byte{1, 31, 0, 8, 0, 0, 0, 0}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestPeakRSSI(t *testing.T) {
	var b, out []byte
	b = PeakRSSI()
	out = []byte{134, 203}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestReaderEventNotificationData(t *testing.T) {
	var b, out []byte
	b = ReaderEventNotificationData(1470125350)
	out = []byte{0, 246, 0, 22, 0, 128, 0, 12, 0, 0, 0, 0, 87, 160, 85, 38, 1, 0, 0, 6, 0, 0}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

/*
func TestTagReportData(t *testing.T) {
	var b, out, dummy []byte
	b = TagReportData(dummy, dummy)
	out = []byte{0, 240}
	if !bytes.Equal(b[:len(out)], out) {
		t.Errorf("%v, want %v", b, out)
	}
	// TODO: might need content length verifications
	t.Skip()
}
*/

func TestUTCTimeStamp(t *testing.T) {
	var b, out []byte
	b = UTCTimeStamp(1470125350)
	out = []byte{0, 128, 0, 12, 0, 0, 0, 0, 87, 160, 85, 38}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestGeneralDeviceCapabilities(t *testing.T) {
	data := GeneralDeviceCapabilities()
	if len(data) != 2+2+2+2+4+4+4+4+4 {
		t.Fatalf("unexpected length: %d", len(data))
	}
	if got := binary.BigEndian.Uint16(data[0:2]); got != 137 {
		t.Fatalf("type = %d, want 137", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 28 {
		t.Fatalf("length = %d, want 28", got)
	}
	if got := binary.BigEndian.Uint16(data[4:6]); got != 52 {
		t.Fatalf("max antennas = %d, want 52", got)
	}
	if got := binary.BigEndian.Uint16(data[6:8]); got != 16384 {
		t.Fatalf("clock support = %d, want 16384", got)
	}
	if got := binary.BigEndian.Uint32(data[8:12]); got != 25882 {
		t.Fatalf("manufacturer = %d, want 25882", got)
	}
	if got := binary.BigEndian.Uint32(data[12:16]); got != 2001007 {
		t.Fatalf("model = %d, want 2001007", got)
	}
}

func TestLLRPCapabilities(t *testing.T) {
	data := LLRPCapabilities()
	if len(data) != 2+2+1+1+2+4+4+4+4+4 {
		t.Fatalf("unexpected length: %d", len(data))
	}
	if got := binary.BigEndian.Uint16(data[0:2]); got != 142 {
		t.Fatalf("type = %d, want 142", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 28 {
		t.Fatalf("length = %d, want 28", got)
	}
	if got := data[4]; got != 72 {
		t.Fatalf("supportedFeatures = %d, want 72", got)
	}
}

func TestRegulatoryCapabilities(t *testing.T) {
	data := RegulatoryCapabilities()
	if got := binary.BigEndian.Uint16(data[0:2]); got != 143 {
		t.Fatalf("type = %d, want 143", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 1189+8 {
		t.Fatalf("length = %d, want 1197", got)
	}
	if got := binary.BigEndian.Uint16(data[4:6]); got != 840 {
		t.Fatalf("country = %d, want 840", got)
	}
	if got := binary.BigEndian.Uint16(data[6:8]); got != 1 {
		t.Fatalf("standards = %d, want 1", got)
	}
	if !bytes.Equal(data[8:], UHFCapabilities(52)) {
		t.Fatalf("UHF capabilities payload mismatch")
	}
}

func TestC1G2LLRPCapabilities(t *testing.T) {
	data := C1G2LLRPCapabilities()
	if got := binary.BigEndian.Uint16(data[0:2]); got != 327 {
		t.Fatalf("type = %d, want 327", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 7 {
		t.Fatalf("length = %d, want 7", got)
	}
	if data[4] != 64 {
		t.Fatalf("flags = %d, want 64", data[4])
	}
	if got := binary.BigEndian.Uint16(data[5:7]); got != 2 {
		t.Fatalf("max filters = %d, want 2", got)
	}
}

func TestAntennaConfigurationIncludesSubparameters(t *testing.T) {
	ac := AntennaConfiguration(5)
	if got := binary.BigEndian.Uint16(ac[0:2]); got != 222 {
		t.Fatalf("type = %d, want 222", got)
	}
	if got := binary.BigEndian.Uint16(ac[4:6]); got != 5 {
		t.Fatalf("antenna id = %d, want 5", got)
	}
	offset := 6
	receiver := RFReceiver()
	if !bytes.Equal(ac[offset:offset+len(receiver)], receiver) {
		t.Fatalf("receiver segment mismatch")
	}
	offset += len(receiver)
	transmitter := RFTransmitter()
	if !bytes.Equal(ac[offset:offset+len(transmitter)], transmitter) {
		t.Fatalf("transmitter segment mismatch")
	}
	offset += len(transmitter)
	inventory := C1G2InventoryCommand()
	if !bytes.Equal(ac[offset:offset+len(inventory)], inventory) {
		t.Fatalf("inventory segment mismatch")
	}
}

func TestRFReceiver(t *testing.T) {
	want := []byte{0x00, 0xdf, 0x00, 0x06, 0x00, 0x01}
	if got := RFReceiver(); !bytes.Equal(got, want) {
		t.Fatalf("RFReceiver() = %v, want %v", got, want)
	}
}

func TestRFTransmitter(t *testing.T) {
	want := []byte{0x00, 0xe0, 0x00, 0x0a, 0x00, 0x01, 0x00, 0x00, 0x00, 0x51}
	if got := RFTransmitter(); !bytes.Equal(got, want) {
		t.Fatalf("RFTransmitter() = %v, want %v", got, want)
	}
}

func TestC1G2RFControl(t *testing.T) {
	data := C1G2RFControl()
	if got := binary.BigEndian.Uint16(data[0:2]); got != 335 {
		t.Fatalf("type = %d, want 335", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 8 {
		t.Fatalf("length = %d, want 8", got)
	}
	if got := binary.BigEndian.Uint16(data[4:6]); got != 1000 {
		t.Fatalf("mode index = %d, want 1000", got)
	}
}

func TestC1G2SingulationControl(t *testing.T) {
	data := C1G2SingulationControl()
	if got := binary.BigEndian.Uint16(data[0:2]); got != 336 {
		t.Fatalf("type = %d, want 336", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 11 {
		t.Fatalf("length = %d, want 11", got)
	}
	if data[4] != 0x40 {
		t.Fatalf("session = %x, want 0x40", data[4])
	}
}

func TestUHFCapabilities(t *testing.T) {
	data := UHFCapabilities(1)
	if got := binary.BigEndian.Uint16(data[0:2]); got != 144 {
		t.Fatalf("type = %d, want 144", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 1189 {
		t.Fatalf("length = %d, want 1189", got)
	}
	entries := data[4:]
	firstEntry := TransmitPowerLevelEntry(1, 1000)
	if !bytes.Equal(entries[:len(firstEntry)], firstEntry) {
		t.Fatalf("first power level entry mismatch")
	}
}

func TestTransmitPowerLevelEntry(t *testing.T) {
	data := TransmitPowerLevelEntry(2, 1250)
	if got := binary.BigEndian.Uint16(data[0:2]); got != 145 {
		t.Fatalf("type = %d, want 145", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 8 {
		t.Fatalf("length = %d, want 8", got)
	}
	if got := binary.BigEndian.Uint16(data[4:6]); got != 2 {
		t.Fatalf("id = %d, want 2", got)
	}
	if got := binary.BigEndian.Uint16(data[6:8]); got != 1250 {
		t.Fatalf("power = %d, want 1250", got)
	}
}

func TestFrequencyInformation(t *testing.T) {
	data := FrequencyInformation()
	if got := binary.BigEndian.Uint16(data[0:2]); got != 146 {
		t.Fatalf("type = %d, want 146", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 213 {
		t.Fatalf("length = %d, want 213", got)
	}
	if data[4] != 1 {
		t.Fatalf("hopping flag = %d, want 1", data[4])
	}
}

func TestFrequencyHopTable(t *testing.T) {
	data := FrequencyHopTable()
	if got := binary.BigEndian.Uint16(data[0:2]); got != 147 {
		t.Fatalf("type = %d, want 147", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 208 {
		t.Fatalf("length = %d, want 208", got)
	}
	if data[4] != 1 {
		t.Fatalf("table id = %d, want 1", data[4])
	}
	if data[5] != 0 {
		t.Fatalf("reserved = %d, want 0", data[5])
	}
	if got := binary.BigEndian.Uint16(data[6:8]); got != 50 {
		t.Fatalf("num hops = %d, want 50", got)
	}
	firstFrequency := binary.BigEndian.Uint32(data[8:12])
	if firstFrequency != 903250 {
		t.Fatalf("first frequency = %d, want 903250", firstFrequency)
	}
}

func TestC1G2UHFModeRFTable(t *testing.T) {
	data := C1G2UHFModeRFTable()
	if got := binary.BigEndian.Uint16(data[0:2]); got != 328 {
		t.Fatalf("type = %d, want 328", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 324 {
		t.Fatalf("length = %d, want 324", got)
	}
	entries := data[4:]
	entrySize := len(C1G2UHFModeRFTableEntry(0))
	if len(entries)%entrySize != 0 {
		t.Fatalf("unexpected entry size")
	}
	first := entries[:entrySize]
	if got := binary.BigEndian.Uint16(first[0:2]); got != 329 {
		t.Fatalf("first entry type = %d, want 329", got)
	}
	if got := binary.BigEndian.Uint16(first[2:4]); got != 32 {
		t.Fatalf("first entry length = %d, want 32", got)
	}
	if got := binary.BigEndian.Uint32(first[4:8]); got != 0 {
		t.Fatalf("first mode id = %d, want 0", got)
	}
	last := entries[len(entries)-entrySize:]
	if got := binary.BigEndian.Uint32(last[4:8]); got != 1004 {
		t.Fatalf("last mode id = %d, want 1004", got)
	}
}

func TestC1G2UHFModeRFTableEntry(t *testing.T) {
	data := C1G2UHFModeRFTableEntry(3)
	if got := binary.BigEndian.Uint16(data[0:2]); got != 329 {
		t.Fatalf("type = %d, want 329", got)
	}
	if got := binary.BigEndian.Uint16(data[2:4]); got != 32 {
		t.Fatalf("length = %d, want 32", got)
	}
	if got := binary.BigEndian.Uint32(data[4:8]); got != 3 {
		t.Fatalf("mode id = %d, want 3", got)
	}
	if data[8] != 80 {
		t.Fatalf("DR flags = %d, want 80", data[8])
	}
}
