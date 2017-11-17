package main

import (
	"reflect"
	"testing"
)

func TestGetFilterValue(t *testing.T) {
	type args struct {
		fv string
	}
	tests := []struct {
		name       string
		args       args
		wantFilter []rune
	}{
		{"000", args{"0"}, []rune("000")},
		{"001", args{"1"}, []rune("001")},
		{"010", args{"2"}, []rune("010")},
		{"011", args{"3"}, []rune("011")},
		{"100", args{"4"}, []rune("100")},
		{"101", args{"5"}, []rune("101")},
		{"110", args{"6"}, []rune("110")},
		{"111", args{"7"}, []rune("111")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFilter := GetFilterValue(tt.args.fv); !reflect.DeepEqual(gotFilter, tt.wantFilter) {
				t.Errorf("GetFilterValue() = %v, want %v", gotFilter, tt.wantFilter)
			}
		})
	}
}

func TestGetIndivisualAssetReference(t *testing.T) {
	type args struct {
		iar     string
		cpSizes []int
	}
	tests := []struct {
		name                         string
		args                         args
		wantIndivisualAssetReference []rune
	}{
		{"5678", args{"5678", []int{24, 7}}, []rune("0000000000000000000000000000000000000000000001011000101110")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndivisualAssetReference := GetIndivisualAssetReference(tt.args.iar, tt.args.cpSizes); !reflect.DeepEqual(gotIndivisualAssetReference, tt.wantIndivisualAssetReference) {
				t.Errorf("GetIndivisualAssetReference() = %v, want %v", string(gotIndivisualAssetReference), string(tt.wantIndivisualAssetReference))
			}
		})
	}
}

func TestGetItemReference(t *testing.T) {
	type args struct {
		ir      string
		cpSizes []int
	}
	tests := []struct {
		name              string
		args              args
		wantItemReference []rune
	}{
		{"1", args{"1", []int{20, 6}}, []rune("000000000000000000000001")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotItemReference := GetItemReference(tt.args.ir, tt.args.cpSizes); !reflect.DeepEqual(gotItemReference, tt.wantItemReference) {
				t.Errorf("GetItemReference() = %v, want %v", gotItemReference, tt.wantItemReference)
			}
		})
	}
}

func TestGetPartitionAndCompanyPrefix(t *testing.T) {
	type args struct {
		cp string
	}
	tests := []struct {
		name              string
		args              args
		wantPartition     []rune
		wantCompanyPrefix []rune
		wantCpSizes       []int
	}{
		{"123456", args{"123456"}, []rune("110"), []rune("00011110001001000000"), []int{20, 6}},
		{"1234567", args{"1234567"}, []rune("101"), []rune("000100101101011010000111"), []int{24, 7}},
		{"12345678", args{"12345678"}, []rune("100"), []rune("000101111000110000101001110"), []int{27, 8}},
		{"123456789", args{"123456789"}, []rune("011"), []rune("000111010110111100110100010101"), []int{30, 9}},
		{"1234567890", args{"1234567890"}, []rune("010"), []rune("0001001001100101100000001011010010"), []int{34, 10}},
		{"12345678901", args{"12345678901"}, []rune("001"), []rune("0001011011111110111000001110000110101"), []int{37, 11}},
		{"123456789012", args{"123456789012"}, []rune("000"), []rune("0001110010111110100110010001101000010100"), []int{40, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPartition, gotCompanyPrefix, gotCpSizes := GetPartitionAndCompanyPrefix(tt.args.cp)
			if !reflect.DeepEqual(gotPartition, tt.wantPartition) {
				t.Errorf("GetPartitionAndCompanyPrefix() gotPartition = %v, want %v", gotPartition, tt.wantPartition)
			}
			if !reflect.DeepEqual(gotCompanyPrefix, tt.wantCompanyPrefix) {
				t.Errorf("GetPartitionAndCompanyPrefix() gotCompanyPrefix = %v, want %v", gotCompanyPrefix, tt.wantCompanyPrefix)
			}
			if !reflect.DeepEqual(gotCpSizes, tt.wantCpSizes) {
				t.Errorf("GetPartitionAndCompanyPrefix() gotCpSizes = %v, want %v", gotCpSizes, tt.wantCpSizes)
			}
		})
	}
}

func TestGetSerial(t *testing.T) {
	type args struct {
		s            string
		serialLength int
	}
	tests := []struct {
		name       string
		args       args
		wantSerial []rune
	}{
		{"1(4) ", args{"1", 4}, []rune("0001")},
		{"10(10)", args{"10", 10}, []rune("0000001010")},
		{"100(38)", args{"100", 38}, []rune("00000000000000000000000000000001100100")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSerial := GetSerial(tt.args.s, tt.args.serialLength); !reflect.DeepEqual(gotSerial, tt.wantSerial) {
				t.Errorf("GetSerial() = %v, want %v", string(gotSerial), string(tt.wantSerial))
			}
		})
	}
}

func TestMakeRuneSliceOfGIAI96(t *testing.T) {
	type args struct {
		cp  string
		fv  string
		iar string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"3474257BF40000000000162E", args{"0614141", "3", "5678"}, []byte{52, 116, 37, 123, 244, 0, 0, 0, 0, 0, 22, 46}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeRuneSliceOfGIAI96(tt.args.cp, tt.args.fv, tt.args.iar)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeRuneSliceOfGIAI96() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRuneSliceOfGIAI96() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeRuneSliceOfGRAI96(t *testing.T) {
	type args struct {
		cp string
		fv string
		at string
		s  string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"3374257BF40C0E400000162E", args{"0614141", "3", "12345", "5678"}, []byte{51, 116, 37, 123, 244, 12, 14, 64, 0, 0, 22, 46}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeRuneSliceOfGRAI96(tt.args.cp, tt.args.fv, tt.args.at, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeRuneSliceOfGRAI96() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRuneSliceOfGRAI96() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeRuneSliceOfSGTIN96(t *testing.T) {
	type args struct {
		cp string
		fv string
		ir string
		s  string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeRuneSliceOfSGTIN96(tt.args.cp, tt.args.fv, tt.args.ir, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeRuneSliceOfSGTIN96() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRuneSliceOfSGTIN96() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeRuneSliceOfSSCC96(t *testing.T) {
	type args struct {
		cp string
		fv string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeRuneSliceOfSSCC96(tt.args.cp, tt.args.fv)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeRuneSliceOfSSCC96() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRuneSliceOfSSCC96() = %v, want %v", got, tt.want)
			}
		})
	}
}
