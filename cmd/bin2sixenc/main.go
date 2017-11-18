package main

import (
	"fmt"
	"os"

	"github.com/iomz/go-llrp/binutil"
)

func bin2sixenc(bs []rune) []rune {
	sixenc := make([]rune, len(bs)/6)
	for i := 0; i*6+6 < len(bs); i++ {
		r, err := binutil.Parse6BinRuneSliceToRune([]rune(bs[i*6 : i*6+6]))
		if err != nil {
			panic(err)
		}
		sixenc[i] = r
	}
	return sixenc
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Must pass a string input")
	}
	fmt.Printf(string(bin2sixenc([]rune(os.Args[1]))))
}