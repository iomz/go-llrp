package llrp

import (
	"bytes"
	"testing"
)

func TestC1G2PC(t *testing.T) {
	var b, out []byte
	var dummy = "3000"
	b = C1G2PC(dummy)
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
	t.Skip()
}

func TestKeepaliveSpec(t *testing.T) {
	var b, out []byte
	b = KeepaliveSpec()
	out = []byte{0, 220, 0, 9, 1, 0, 0, 39, 16}
	if !bytes.Equal(b, out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestLLRPStatus(t *testing.T) {
	var b, out []byte
	b = LLRPStatus()
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
	b = ReaderEventNotificationData()
	out = []byte{0, 246, 0, 22, 0, 128, 0, 12, 0, 0, 0, 0, 0}
	if !bytes.Equal(b[:len(out)], out) {
		t.Errorf("%v, want %v", b, out)
	}
}

func TestTagReportData(t *testing.T) {
	var b, out, dummy []byte
	b = TagReportData(dummy, dummy, dummy, dummy)
	out = []byte{0, 240}
	if !bytes.Equal(b[:len(out)], out) {
		t.Errorf("%v, want %v", b, out)
	}
	// TODO: might need content length verifications
	t.Skip()
}

func TestUTCTimeStamp(t *testing.T) {
	var b, out []byte
	b = UTCTimeStamp()
	out = []byte{0, 128, 0, 12}
	if !bytes.Equal(b[:len(out)], out) {
		t.Errorf("%v, want %v", b, out)
	}
	// TODO: might need content length verifications
	t.Skip()
}
