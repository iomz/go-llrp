package main

import (
	"fmt"
	"os"

	"github.com/iomz/go-llrp/binutil"
)

func sixenc2bin(sixenc []rune) []rune {
	var bs []rune
	for i := 0; i < len(sixenc); i++ {
		r := binutil.ParseRuneTo6BinRuneSlice(sixenc[i])
		bs = append(bs, r...)
	}
	return bs
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Must pass a string input")
	}
	fmt.Printf(string(sixenc2bin([]rune(os.Args[1]))))
}
