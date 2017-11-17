package main

import (
	"reflect"
	"testing"
)

func Test_sixenc2bin(t *testing.T) {
	type args struct {
		sixenc []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"7BABCD1234567", args{[]rune("7BABCD1234567")}, []rune("110111000010000001000010000011000100110001110010110011110100110101110110110111")},
		{"1JLA104623JRB1410-08", args{[]rune("1JLA104623JRB1410-08")}, []rune("110001001010001100000001110001110000110100110110110010110011001010010010000010110001110100110001110000101101110000111000")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sixenc2bin(tt.args.sixenc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sixenc2bin() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
