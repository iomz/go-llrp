package binutil

import (
	"reflect"
	"testing"
)

func TestGenerateNLengthAlphabetString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"n = 0", args{0}, ""},
		{"n = 1", args{1}, "A"},
		{"n = 2", args{2}, "AB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateNLengthHexString(tt.args.n); len(got) != len(tt.want) {
				t.Errorf("GenerateNLengthHexString() = %v, want %v length", got, len(tt.want))
			}
		})
	}
}

func TestGenerateNLengthDigitString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"n = 0", args{0}, ""},
		{"n = 1", args{1}, "0"},
		{"n = 2", args{2}, "01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateNLengthHexString(tt.args.n); len(got) != len(tt.want) {
				t.Errorf("GenerateNLengthHexString() = %v, want %v length", got, len(tt.want))
			}
		})
	}
}

func TestGenerateNLengthHexString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"n = 0", args{0}, ""},
		{"n = 1", args{1}, "a"},
		{"n = 2", args{2}, "ab"},
		{"n = 32", args{32}, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
		{"n = 64", args{64}, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateNLengthHexString(tt.args.n); len(got) != len(tt.want) {
				t.Errorf("GenerateNLengthHexString() = %v, want %v length", got, len(tt.want))
			}
		})
	}
}

func TestGenerateNLengthRandomBinRuneSlice(t *testing.T) {
	type args struct {
		n   int
		max uint
	}
	tests := []struct {
		name  string
		args  args
		want  []rune
		want1 uint
	}{
		{"n = 0", args{0, 0}, []rune(""), 0},
		{"n = 2", args{2, 0}, []rune("11"), 3},
		{"n = 64, max = 16", args{64, 16}, []rune("0000000000000000000000000000000000000000000000000000000000010000"), 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GenerateNLengthRandomBinRuneSlice(tt.args.n, tt.args.max)
			if len(got) != len(tt.want) {
				t.Errorf("GenerateNLengthHexString() = %v, want %v length", got, len(tt.want))
			}
			if got1 > tt.want1 {
				t.Errorf("GenerateNLengthRandomBinRuneSlice() got1 = %v, want less than %v", got1, tt.want1)
			}
		})
	}
}

func TestGenerateNLengthZeroPaddingRuneSlice(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"n = 0", args{0}, []rune("")},
		{"n = 2", args{2}, []rune("00")},
		{"n = 32", args{32}, []rune("00000000000000000000000000000000")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateNLengthZeroPaddingRuneSlice(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateNLengthZeroPaddingRuneSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{"min = 0, max = 100", args{0, 100}},
		{"min = 35, max = 40", args{35, 40}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandomInt(tt.args.min, tt.args.max)
			if got < tt.args.min {
				t.Errorf("GenerateRandomInt() = %v, want > %v", got, tt.args.min)
			} else if got > tt.args.max {
				t.Errorf("GenerateRandomInt() = %v, want < %v", got, tt.args.max)
			}
		})
	}
}

func TestPack(t *testing.T) {
	type args struct {
		data []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pack(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse6BinRuneSliceToRune(t *testing.T) {
	type args struct {
		r []rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"101000 -> '('", args{[]rune("101000")}, '('},
		{"110000 -> '0'", args{[]rune("110000")}, '0'},
		{"000001 -> '!'", args{[]rune("000001")}, 'A'},
		{"111111 -> '?'", args{[]rune("111111")}, '?'},
		{"000000 -> '@'", args{[]rune("000000")}, '@'},
		{"011101 -> ']'", args{[]rune("011101")}, ']'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Parse6BinRuneSliceToRune(tt.args.r); got != tt.want {
				t.Errorf("Parse6BinRuneSliceToRune() = %c, want %c", got, tt.want)
			} else if err != nil {
				t.Errorf(err.Error())
			}
		})
	}
}

func TestParseBinRuneSliceToUint8Slice(t *testing.T) {
	type args struct {
		bs []rune
	}
	tests := []struct {
		name    string
		args    args
		want    []uint8
		wantErr bool
	}{
		{"01110101 -> 117", args{[]rune("01110101")}, []uint8{117}, false},
		{"0000000011111111 -> 0,255", args{[]rune("0000000011111111")}, []uint8{0, 255}, false},
		{" -> error", args{[]rune("")}, nil, true},
		{"0000 -> error", args{[]rune("0000")}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBinRuneSliceToUint8Slice(tt.args.bs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBinRuneSliceToUint8Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseBinRuneSliceToUint8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDecimalStringToBinRuneSlice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"123456789 -> 1001001100101100000001011010010", args{"1234567890"}, []rune("1001001100101100000001011010010")},
		{"4294967296 -> 100000000000000000000000000000000", args{"4294967296"}, []rune("100000000000000000000000000000000")},
		{"1 -> 1", args{"1"}, []rune("1")},
		{"2 -> 10", args{"2"}, []rune("10")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDecimalStringToBinRuneSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDecimalStringToBinRuneSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseHexStringToBinString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"0123456789ABCDEF -> 000000010010001101000101011001111000100110101011110011011111", args{"0123456789ABCDEF"}, "0000000100100011010001010110011110001001101010111100110111101111", false},
		{"string -> error", args{"string"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseHexStringToBinString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseHexStringToBinString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseHexStringToBinString() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestParseRuneTo6BinRuneSlice(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"r = '('", args{'('}, []rune("101000")},
		{"r = '0'", args{'0'}, []rune("110000")},
		{"r = 'A'", args{'A'}, []rune("000001")},
		{"r = '?'", args{'?'}, []rune("111111")},
		{"r = '@'", args{'@'}, []rune("000000")},
		{"r = ']'", args{']'}, []rune("011101")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseRuneTo6BinRuneSlice(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRuneTo6BinRuneSlice() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
