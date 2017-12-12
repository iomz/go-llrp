package main

import (
	"reflect"
	"testing"
)

func TestGetAssetType(t *testing.T) {
	type args struct {
		at string
		pr map[PartitionTableKey]int
	}
	tests := []struct {
		name          string
		args          args
		wantAssetType []rune
	}{
		{"5678", args{"5678", GRAI96PartitionTable[7]}, []rune("00000001011000101110")},
		{"(blank)", args{"", GRAI96PartitionTable[7]}, []rune("00000001011000101110")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAssetType := GetAssetType(tt.args.at, tt.args.pr); !reflect.DeepEqual(gotAssetType, tt.wantAssetType) {
				if tt.args.at == "" && len(gotAssetType) == tt.args.pr[ATBits] {
					// pass
				} else {
					t.Errorf("GetAssetType() = %v, want %v", string(gotAssetType), string(tt.wantAssetType))
				}
			}
		})
	}
}

func TestGetCompanyPrefix(t *testing.T) {
	type args struct {
		cp string
		pt map[int]map[PartitionTableKey]int
	}
	tests := []struct {
		name              string
		args              args
		wantCompanyPrefix []rune
	}{
		{"123456", args{"123456", SGTIN96PartitionTable}, []rune("00011110001001000000")},
		{"1234567", args{"1234567", SGTIN96PartitionTable}, []rune("000100101101011010000111")},
		{"12345678", args{"12345678", SGTIN96PartitionTable}, []rune("000101111000110000101001110")},
		{"123456789", args{"123456789", SGTIN96PartitionTable}, []rune("000111010110111100110100010101")},
		{"1234567890", args{"1234567890", SGTIN96PartitionTable}, []rune("0001001001100101100000001011010010")},
		{"12345678901", args{"12345678901", SGTIN96PartitionTable}, []rune("0001011011111110111000001110000110101")},
		{"123456789012", args{"123456789012", SGTIN96PartitionTable}, []rune("0001110010111110100110010001101000010100")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCompanyPrefix := GetCompanyPrefix(tt.args.cp, tt.args.pt); !reflect.DeepEqual(gotCompanyPrefix, tt.wantCompanyPrefix) {
				t.Errorf("GetCompanyPrefix() = %v, want %v", gotCompanyPrefix, tt.wantCompanyPrefix)
			}
		})
	}
}

func TestGetExtension(t *testing.T) {
	type args struct {
		e  string
		pr map[PartitionTableKey]int
	}
	tests := []struct {
		name          string
		args          args
		wantExtension []rune
	}{
		{"1234567890", args{"1234567890", SSCC96PartitionTable[7]}, []rune("0001001001100101100000001011010010")},
		{"(blank)", args{"", SSCC96PartitionTable[7]}, []rune("0001001001100101100000001011010010")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExtension := GetExtension(tt.args.e, tt.args.pr); !reflect.DeepEqual(gotExtension, tt.wantExtension) {
				if tt.args.e == "" && len(gotExtension) == tt.args.pr[EBits] {
					// pass
				} else {
					t.Errorf("GetExtension() = %v, want %v", string(gotExtension), string(tt.wantExtension))
				}
			}
		})
	}
}

func TestGetFilter(t *testing.T) {
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
		{"(blank)", args{""}, []rune("111")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFilter := GetFilter(tt.args.fv); !reflect.DeepEqual(gotFilter, tt.wantFilter) {
				if tt.args.fv == "" && len(gotFilter) == 3 {
					// pass
				} else {
					t.Errorf("GetFilter() = %v, want %v", gotFilter, tt.wantFilter)
				}
			}
		})
	}
}

func TestGetIndivisualAssetReference(t *testing.T) {
	type args struct {
		iar string
		pr  map[PartitionTableKey]int
	}
	tests := []struct {
		name                         string
		args                         args
		wantIndivisualAssetReference []rune
	}{
		{"5678", args{"5678", GIAI96PartitionTable[7]}, []rune("0000000000000000000000000000000000000000000001011000101110")},
		{"(blank)", args{"", GIAI96PartitionTable[7]}, []rune{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndivisualAssetReference := GetIndivisualAssetReference(tt.args.iar, tt.args.pr); !reflect.DeepEqual(gotIndivisualAssetReference, tt.wantIndivisualAssetReference) {
				if tt.args.iar == "" && tt.args.pr[IARBits] == len(gotIndivisualAssetReference) {
					// pass
				} else {
					t.Errorf("GetIndivisualAssetReference() = %v, want %v", gotIndivisualAssetReference, tt.wantIndivisualAssetReference)
				}
			}
		})
	}
}

func TestGetItemReference(t *testing.T) {
	type args struct {
		ir string
		pr map[PartitionTableKey]int
	}
	tests := []struct {
		name              string
		args              args
		wantItemReference []rune
	}{
		{"1", args{"1", map[PartitionTableKey]int{IRBits: 20, IRDigits: 6}}, []rune("00000000000000000001")},
		{"(blank)", args{"", map[PartitionTableKey]int{IRBits: 20, IRDigits: 6}}, []rune("00000000000000000001")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotItemReference := GetItemReference(tt.args.ir, tt.args.pr); !reflect.DeepEqual(gotItemReference, tt.wantItemReference) {
				t.Logf("GetItemReference() = %v, want %v", string(gotItemReference), string(tt.wantItemReference))
				if len(gotItemReference) != len(tt.wantItemReference) {
					t.Errorf("len(GetItemReference()) = %v, want %v", len(string(gotItemReference)), len(string(tt.wantItemReference)))
				}
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
		{"(blank)(38)", args{"", 38}, []rune("00000000000000000000000000000001100100")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSerial := GetSerial(tt.args.s, tt.args.serialLength); !reflect.DeepEqual(gotSerial, tt.wantSerial) {
				t.Logf("GetSerial() = %v, want %v", string(gotSerial), string(tt.wantSerial))
				if len(gotSerial) != len(tt.wantSerial) {
					t.Errorf("len(GetSerial()) = %v, want %v", len(string(gotSerial)), len(string(tt.wantSerial)))
				}
			}
		})
	}
}

func TestMakeGIAI96(t *testing.T) {
	type args struct {
		pf  bool
		fv  string
		cp  string
		iar string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   string
		want2   string
		wantErr bool
	}{
		{"3474257BF40000000000162E", args{false, "3", "0614141", "5678"}, []byte{52, 116, 37, 123, 244, 0, 0, 0, 0, 0, 22, 46}, "", "", false},
		{"3474257BF40000000000162E", args{true, "3", "0614141", "5678"}, []byte{}, "001101000111010000100101011110111111010000000000000000000000000000000000000000000001011000101110", "GIAI-96_3_5_0614141_5678", false},
		{"3474257BF40000000000162E", args{true, "3", "0614141", ""}, []byte{}, "00110100011101000010010101111011111101", "GIAI-96_3_5_0614141", false},
		{"3474257BF40000000000162E", args{true, "3", "", ""}, []byte{}, "00110100011", "GIAI-96_3", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := MakeGIAI96(tt.args.pf, tt.args.fv, tt.args.cp, tt.args.iar)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeGIAI96() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeGIAI96() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeGIAI96() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("MakeGIAI96() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestMakeGRAI96(t *testing.T) {
	type args struct {
		pf  bool
		fv  string
		cp  string
		at  string
		ser string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   string
		want2   string
		wantErr bool
	}{
		{"3374257BF40C0E400000162E", args{false, "3", "0614141", "12345", "5678"}, []byte{51, 116, 37, 123, 244, 12, 14, 64, 0, 0, 22, 46}, "", "", false},
		{"3374257BF40C0E400000162E", args{true, "3", "0614141", "12345", "5678"}, []byte{}, "001100110111010000100101011110111111010000001100000011100100000000000000000000000001011000101110", "GRAI-96_3_5_0614141_12345_5678", false},
		{"3374257BF40C0E400000162E", args{true, "3", "0614141", "12345", ""}, []byte{}, "0011001101110100001001010111101111110100000011000000111001", "GRAI-96_3_5_0614141_12345", false},
		{"3374257BF40C0E400000162E", args{true, "3", "0614141", "", ""}, []byte{}, "00110011011101000010010101111011111101", "GRAI-96_3_5_0614141", false},
		{"3374257BF40C0E400000162E", args{true, "3", "", "", ""}, []byte{}, "00110011011", "GRAI-96_3", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := MakeGRAI96(tt.args.pf, tt.args.fv, tt.args.cp, tt.args.at, tt.args.ser)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeGRAI96() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeGRAI96() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeGRAI96() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("MakeGRAI96() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestMakeSGTIN96(t *testing.T) {
	type args struct {
		pf  bool
		fv  string
		cp  string
		ir  string
		ser string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   string
		want2   string
		wantErr bool
	}{
		{"3074257BF7194E4000001A85", args{false, "3", "0614141", "812345", "6789"}, []byte{48, 116, 37, 123, 247, 25, 78, 64, 0, 0, 26, 133}, "", "", false},
		{"3074257BF7194E4000001A85", args{true, "3", "0614141", "812345", "6789"}, []byte{}, "001100000111010000100101011110111111011100011001010011100100000000000000000000000001101010000101", "SGTIN-96_3_5_0614141_812345_6789", false},
		{"3074257BF7194E4000001A85", args{true, "3", "0614141", "812345", ""}, []byte{}, "0011000001110100001001010111101111110111000110010100111001", "SGTIN-96_3_5_0614141_812345", false},
		{"3074257BF7194E4000001A85", args{true, "3", "0614141", "", ""}, []byte{}, "00110000011101000010010101111011111101", "SGTIN-96_3_5_0614141", false},
		{"3074257BF7194E4000001A85", args{true, "3", "", "", ""}, []byte{}, "00110000011", "SGTIN-96_3", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := MakeSGTIN96(tt.args.pf, tt.args.fv, tt.args.cp, tt.args.ir, tt.args.ser)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeSGTIN96() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeSGTIN96() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeSGTIN96() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("MakeSGTIN96() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestMakeSSCC96(t *testing.T) {
	type args struct {
		pf  bool
		fv  string
		cp  string
		ext string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   string
		want2   string
		wantErr bool
	}{
		{"3174257BF4499602D2000000", args{false, "3", "0614141", "1234567890"}, []byte{49, 116, 37, 123, 244, 73, 150, 2, 210, 0, 0, 0}, "", "", false},
		{"3174257BF4499602D2000000", args{true, "3", "0614141", "1234567890"}, []byte{}, "001100010111010000100101011110111111010001001001100101100000001011010010", "SSCC-96_3_5_0614141_1234567890", false},
		{"3174257BF4499602D2000000", args{true, "3", "0614141", ""}, []byte{}, "00110001011101000010010101111011111101", "SSCC-96_3_5_0614141", false},
		{"3174257BF4499602D2000000", args{true, "3", "", ""}, []byte{}, "00110001011", "SSCC-96_3", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := MakeSSCC96(tt.args.pf, tt.args.fv, tt.args.cp, tt.args.ext)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeSSCC96() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeSSCC96() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeSSCC96() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("MakeSSCC96() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
