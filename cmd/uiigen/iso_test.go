package main

import (
	"reflect"
	"testing"
)

func TestGetISO6346CD(t *testing.T) {
	type args struct {
		cn string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"CSQU305438", args{"CSQU305438"}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetISO6346CD(tt.args.cn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetISO6346CD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetISO6346CD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeISO17363(t *testing.T) {
	type args struct {
		oc  string
		ei  string
		csn string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   int
		wantErr bool
	}{
		{"A97BCSQU3054383", args{"CSQ", "U", "305438"}, []byte{220, 32, 211, 69, 92, 240, 215, 76, 248, 206}, 80, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := MakeISO17363(tt.args.oc, tt.args.ei, tt.args.csn)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeISO17363() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeISO17363() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeISO17363() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMakeISO17365(t *testing.T) {
	type args struct {
		di  string
		iac string
		cin string
		sn  string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   int
		wantErr bool
	}{
		{"25SUN043325711MH8031200000000001", args{"25S", "UN", "043325711", "MH8031200000000001"}, []byte{203, 84, 213, 59, 13, 51, 207, 45, 119, 199, 19, 72, 227, 12, 241, 203, 12, 48, 195, 12, 48, 195, 12, 49}, 192, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := MakeISO17365(tt.args.di, tt.args.iac, tt.args.cin, tt.args.sn)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeISO17365() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeISO17365() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeISO17365() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPad6BitEncodingRuneSlice(t *testing.T) {
	type args struct {
		bs []rune
	}
	tests := []struct {
		name  string
		args  args
		want  []rune
		want1 int
	}{
		{"0000", args{[]rune("0000")}, []rune("0000100000100000"), 16},
		{"0000000000000000", args{[]rune("0000000000000000")}, []rune("0000000000000000"), 16},
		//{"", args{""}, []rune{""}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Pad6BitEncodingRuneSlice(tt.args.bs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pad6BitEncodingRuneSlice() got = %v, want %v", string(got), string(tt.want))
			}
			if got1 != tt.want1 {
				t.Errorf("Pad6BitEncodingRuneSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
