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

func TestMakeRuneSliceOfISO17363(t *testing.T) {
	type args struct {
		afi string
		oc  string
		ei  string
		csn string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 int
	}{
		{"A97BCSQU3054383", args{"A9", "CSQ", "U", "305438"}, []byte{169, 220, 32, 211, 69, 92, 240, 215, 76, 248, 206, 8}, 96},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MakeRuneSliceOfISO17363(tt.args.afi, tt.args.oc, tt.args.ei, tt.args.csn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRuneSliceOfISO17363() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeRuneSliceOfISO17363() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMakeRuneSliceOfISO17365(t *testing.T) {
	type args struct {
		afi string
		di  string
		iac string
		cin string
		sn  string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 int
	}{
		{"", args{"A1", "25S", "UN", "043325711", "MH8031200000000001"}, []byte{161, 203, 84, 213, 59, 13, 51, 207, 45, 119, 199, 19, 72, 227, 12, 241, 203, 12, 48, 195, 12, 48, 195, 12, 49, 130}, 208},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MakeRuneSliceOfISO17365(tt.args.afi, tt.args.di, tt.args.iac, tt.args.cin, tt.args.sn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRuneSliceOfISO17365() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeRuneSliceOfISO17365() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
