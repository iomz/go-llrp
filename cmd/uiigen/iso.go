// This file contains ISO binary encoding scheme
package main

import (
	"fmt"
	"os"

	"github.com/iomz/go-llrp/binutil"
)
// MakeRuneSliceOfISO17365 generates a random 17367 code
func MakeRuneSliceOfISO17365(afi string, di string, iac string, cin string, sn string) ([]byte, int) {
	applicationFamilyIdentifier, _ := binutil.ParseHexStringToBinString(afi)
	dataIdentifier := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(di))
	issuingAgencyCode := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(iac))
	companyIdentification := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(cin))
	serialNumber := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(sn))

	var bs []rune
	bs = append(bs, []rune(applicationFamilyIdentifier)...)
	bs = append(bs, dataIdentifier...)
	bs = append(bs, issuingAgencyCode...)
	bs = append(bs, companyIdentification...)
	bs = append(bs, serialNumber...)

	length := len(applicationFamilyIdentifier)+len(dataIdentifier)+len(issuingAgencyCode)+len(companyIdentification)+len(serialNumber)
	remainder := length % 16
	var padding []rune
	if remainder != 0 {
		padRuneSlice := binutil.ParseDecimalStringToBinRuneSlice("32") // pad string "100000"
		for i:=0; i<16-remainder; i++ {
			padding = append(padding, padRuneSlice[i%6])
		}
		bs = append(bs, padding...)
		length += 16-remainder
	}

	p, err := binutil.ParseBinRuneSliceToUint8Slice(bs)
	if err != nil {
		fmt.Println("Something went wrong!")
		fmt.Println(err)
		os.Exit(1)
	}

	var iso17365 = []interface{}{p}

	return binutil.Pack(iso17365), length
}
