package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_parseArg(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3", args{[]string{"randomdigit", "3"}}, 3},
		{"100", args{[]string{"randomdigit", "100"}}, 100},
		{"001", args{[]string{"randomdigit", "001"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseArg(tt.args.args); got != tt.want {
				t.Errorf("parseArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printAlphabetString(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{"3", args{3}, "ABC\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			printAlphabetString(w, tt.args.i)
			if gotW := w.String(); len(gotW) != len(tt.wantW) {
				t.Errorf("printAlphabetString() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
