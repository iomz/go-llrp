package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"unicode"
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

func Test_parseArgPanicsOnInsufficientArgs(t *testing.T) {
	assertPanicsWith(t, "insufficient arg", func() {
		parseArg([]string{"randomalphabet"})
	})
}

func Test_parseArgPanicsOnInvalidNumber(t *testing.T) {
	assertPanicsWith(t, "invalid syntax", func() {
		parseArg([]string{"randomalphabet", "not-a-number"})
	})
}

func Test_printAlphabetString(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		args   args
		length int
	}{
		{"3", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			printAlphabetString(w, tt.args.i)
			gotW := w.String()
			if !strings.HasSuffix(gotW, "\n") {
				t.Fatalf("expected trailing newline, got %q", gotW)
			}
			body := strings.TrimSuffix(gotW, "\n")
			if len(body) != tt.length {
				t.Fatalf("expected body length %d, got %d", tt.length, len(body))
			}
			for _, r := range body {
				if !unicode.IsUpper(r) || !unicode.IsLetter(r) {
					t.Fatalf("unexpected rune %q in %q", r, body)
				}
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func assertPanicsWith(t *testing.T, substr string, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic but none occurred")
		} else if substr != "" {
			switch v := r.(type) {
			case error:
				if !strings.Contains(v.Error(), substr) {
					t.Fatalf("panic error %q does not contain %q", v.Error(), substr)
				}
			case string:
				if !strings.Contains(v, substr) {
					t.Fatalf("panic string %q does not contain %q", v, substr)
				}
			default:
				if !strings.Contains(fmt.Sprint(v), substr) {
					t.Fatalf("panic value %v does not contain %q", v, substr)
				}
			}
		}
	}()
	fn()
}
