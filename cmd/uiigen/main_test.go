package main

import (
	"os"
	"testing"
)

func TestCheckIfStringInSlice(t *testing.T) {
	type args struct {
		a    string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIfStringInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("CheckIfStringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
