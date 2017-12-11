package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"github.com/iomz/go-llrp/binutil"
)

func parseArg(args []string) int {
	if len(args) < 2 {
		panic("insufficient arg")
	}
	i, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	return i
}

func printHexString(w io.Writer, i int) {
	fmt.Fprintf(w, "%s\n", binutil.GenerateNLengthHexString(i))
}

func main() {
	printHexString(os.Stdout, parseArg(os.Args))
}
