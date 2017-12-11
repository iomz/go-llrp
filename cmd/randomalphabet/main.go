package main

import (
	"fmt"
	"github.com/iomz/go-llrp/binutil"
	"io"
	"os"
	"strconv"
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

func printAlphabetString(w io.Writer, i int) {
	fmt.Fprintf(w, "%s\n", binutil.GenerateNLengthAlphabetString(i))
}

func main() {
	printAlphabetString(os.Stdout, parseArg(os.Args))
}
