package llrp

import (
	"bytes"
	"testing"
)

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
