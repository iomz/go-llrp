package llrp

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func parseMessage(t *testing.T, msg []byte) (uint16, uint32, uint32, []byte) {
	t.Helper()
	if len(msg) < 10 {
		t.Fatalf("message too short: %d bytes", len(msg))
	}
	header := binary.BigEndian.Uint16(msg[0:2])
	length := binary.BigEndian.Uint32(msg[2:6])
	messageID := binary.BigEndian.Uint32(msg[6:10])
	return header, length, messageID, msg[10:]
}

func TestKeepalive(t *testing.T) {
	var b, out []byte
	b = Keepalive(0)
	out = []byte{4, 62, 0, 0, 0, 10, 0, 0, 0, 0}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestKeepaliveAck(t *testing.T) {
	var b, out []byte
	b = KeepaliveAck(0)
	out = []byte{4, 72, 0, 0, 0, 10, 0, 0, 0, 0}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

/*
func TestROAccessReport(t *testing.T) {
	var b, out, dummy []byte
	b = ROAccessReport(dummy, 1000)
	out = []byte{4, 61, 0, 0, 0, 10, 0, 0, 3, 232}
	if !bytes.Equal(b[:len(out)], out) {
		t.Errorf("%v, want %v", b, out)
	}
	// TODO: might need content length verifications
	t.Skip()
}
*/

func TestReaderEventNotification(t *testing.T) {
	var b, out []byte
	b = ReaderEventNotification(1000, 1470125350)
	out = []byte{4, 63, 0, 0, 0, 32, 0, 0, 3, 232}
	if !bytes.Equal(b[:len(out)], out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestSetReaderConfig(t *testing.T) {
	var b, out []byte
	b = SetReaderConfig(1000)
	out = []byte{4, 3, 0, 0, 0, 20, 0, 0, 3, 232, 0, 0, 220, 0, 9, 1, 0, 0, 39, 16}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestSetReaderConfigResponse(t *testing.T) {
	var b, out []byte
	b = SetReaderConfigResponse(1000)
	out = []byte{4, 13, 0, 0, 0, 18, 0, 0, 3, 232, 1, 31, 0, 8, 0, 0, 0, 0}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestGetReaderCapability(t *testing.T) {
	const msgID = 0xdeadbeef
	header, length, id, body := parseMessage(t, GetReaderCapability(msgID))
	if header != uint16(GetReaderCapabilityHeader) {
		t.Fatalf("header = %d, want %d", header, GetReaderCapabilityHeader)
	}
	if id != msgID {
		t.Fatalf("message id = %d, want %d", id, msgID)
	}
	if length != 11 {
		t.Fatalf("length = %d, want 11", length)
	}
	if len(body) != 1 || body[0] != 0 {
		t.Fatalf("body = %v, want single zero byte", body)
	}
}

func TestGetReaderCapabilityResponse(t *testing.T) {
	header, length, id, body := parseMessage(t, GetReaderCapabilityResponse(42))
	if header != uint16(GetReaderCapabilityResponseHeader) {
		t.Fatalf("header = %d, want %d", header, GetReaderCapabilityResponseHeader)
	}
	if id != 42 {
		t.Fatalf("message id = %d, want 42", id)
	}
	expected := append([]byte{}, Status()...)
	expected = append(expected, GeneralDeviceCapabilities()...)
	expected = append(expected, LlrpCapabilities()...)
	expected = append(expected, ReguCapabilities()...)
	expected = append(expected, C1G2llrpCapabilities()...)
	if uint32(len(expected))+10 != length {
		t.Fatalf("reported length %d does not match body size %d", length, len(expected)+10)
	}
	if !bytes.Equal(body, expected) {
		t.Fatalf("body mismatch")
	}
}

func TestGetReaderConfigResponse(t *testing.T) {
	header, length, id, body := parseMessage(t, GetReaderConfigResponse(7))
	if header != uint16(GetReaderConfigResponseHeader) {
		t.Fatalf("header = %d, want %d", header, GetReaderConfigResponseHeader)
	}
	if id != 7 {
		t.Fatalf("message id = %d, want 7", id)
	}
	expected := append([]byte{}, Status()...)
	expected = append(expected, GetReaderConfigResponseIdentification()...)
	if uint32(len(expected))+10 != length {
		t.Fatalf("reported length %d does not match body size %d", length, len(expected)+10)
	}
	if !bytes.Equal(body, expected) {
		t.Fatalf("body mismatch")
	}
}

func TestDeleteResponsesIncludeStatus(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(uint32) []byte
		expected uint16
	}{
		{"DeleteAccessSpecResponse", DeleteAccessSpecResponse, uint16(DeleteAccessSpecResponseHeader)},
		{"DeleteRospecResponse", DeleteRospecResponse, uint16(DeleteRospecResponseHeader)},
		{"AddRospecResponse", AddRospecResponse, uint16(AddRospecResponseHeader)},
		{"EnableRospecResponse", EnableRospecResponse, uint16(EnableRospecResponseHeader)},
	}
	wantStatus := Status()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, length, id, body := parseMessage(t, tt.fn(123))
			if header != tt.expected {
				t.Fatalf("header = %d, want %d", header, tt.expected)
			}
			if id != 123 {
				t.Fatalf("message id = %d, want 123", id)
			}
			if length != uint32(len(wantStatus))+10 {
				t.Fatalf("length = %d, want %d", length, len(wantStatus)+10)
			}
			if !bytes.Equal(body, wantStatus) {
				t.Fatalf("body mismatch")
			}
		})
	}
}

func TestReceiveSensitivityEntries(t *testing.T) {
	const antennas = 4
	entries := ReceiveSensitivityEntries(antennas)
	if len(entries) != antennas {
		t.Fatalf("got %d entries, want %d", len(entries), antennas)
	}
	for i, entry := range entries {
		want := ReceiveSensitivityEntry(uint16(i + 1))
		got, ok := entry.([]byte)
		if !ok {
			t.Fatalf("entry %d is %T, want []byte", i, entry)
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("entry %d mismatch", i)
		}
	}
}

func TestGPIOCapabilities(t *testing.T) {
	got := GPIOCapabilities()
	want := []byte{0x00, 0x8d, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00}
	if !bytes.Equal(got, want) {
		t.Fatalf("GPIOCapabilities() = %v, want %v", got, want)
	}
}

func TestAntennaAirPortList(t *testing.T) {
	list := AntennaAirPortList(2)
	if len(list) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(list))
	}
	for i, raw := range list {
		got, ok := raw.([]byte)
		if !ok {
			t.Fatalf("entry %d has type %T", i, raw)
		}
		want := AntennaAirPort(uint16(i + 1))
		if !bytes.Equal(got, want) {
			t.Fatalf("entry %d mismatch", i)
		}
	}
}

func TestImpinjEnableCutomMessage(t *testing.T) {
	header, length, id, body := parseMessage(t, ImpinjEnableCutomMessage(77))
	if header != uint16(ImpinjEnableCutomMessageHeader) {
		t.Fatalf("header = %d, want %d", header, ImpinjEnableCutomMessageHeader)
	}
	if id != 77 {
		t.Fatalf("message id = %d, want 77", id)
	}
	if length != 23 {
		t.Fatalf("length = %d, want 23", length)
	}
	if len(body) != 13 {
		t.Fatalf("body length = %d, want 13", len(body))
	}
	if binary.BigEndian.Uint32(body[0:4]) != 25822 {
		t.Fatalf("vendor id = %d, want 25822", binary.BigEndian.Uint32(body[0:4]))
	}
	if body[4] != 22 {
		t.Fatalf("subtype = %d, want 22", body[4])
	}
	if !bytes.Equal(body[5:], Status()) {
		t.Fatalf("status payload mismatch")
	}
}
