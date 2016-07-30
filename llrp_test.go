package llrp

import (
	"bytes"
	"errors"
	"testing"
)

func TestCheck(t *testing.T) {
	e := errors.New("dummy error")
	check(nil)
	assertCheckPanic(t, check, e)
}

func assertCheckPanic(t *testing.T, f func(error), e error) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f(e)
}

var packtests = []struct {
	in  []interface{}
	out []byte
}{
	{[]interface{}{uint16(349), uint16(11), uint8(0)}, []byte{1, 93, 0, 11, 0}},
	{[]interface{}{uint8(12), uint8(11), uint32(433)}, []byte{12, 11, 0, 0, 1, 177}},
}

func TestPack(t *testing.T) {
	var b []byte
	for _, tt := range packtests {
		b = Pack(tt.in)
		if !bytes.Equal(b, tt.out) {
			t.Errorf("%v => %v, want %v", tt.in, b, tt.out)
		}
	}
}
