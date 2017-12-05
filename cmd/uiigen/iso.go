// This file contains ISO binary encoding scheme
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/iomz/go-llrp/binutil"
)

// GetISO6346CD returns check digit for container serial number
func GetISO6346CD(cn string) (int, error) {
	if len(cn) != 10 {
		return 0, errors.New("Invalid ISO6346 code provided!")
	}
	cn = strings.ToUpper(cn)
	n := 0.0
	d := 0.5
	for i := 0; i < 10; i++ {
		d *= 2
		n += d * float64(strings.Index("0123456789A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ", string(cn[i])))
	}
	return (int(n) - int(n/11)*11) % 10, nil
}

// MakeRuneSliceOfISO17363 generates a random 17363 code
func MakeRuneSliceOfISO17363(oc string, ei string, csn string) ([]byte, int) {
	di := "7B"
	dataIdentifier := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(di))
	if oc == "" {
		oc = binutil.GenerateNLengthAlphabetString(3)
	}
	ownerCode := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(oc))
	equipmentIdentifier := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(ei))
	if csn == "" {
		csn = binutil.GenerateNLengthDigitString(6)
	} else if 6 > len(csn) {
		leftPadding := binutil.GenerateNLengthZeroPaddingRuneSlice(6 - len(csn))
		csn = string(leftPadding) + csn
	} else if 6 < len(csn) {
		panic("Invalid csn: " + csn)
	}
	cd, err := GetISO6346CD(oc + ei + csn)
	if err != nil {
		panic(err)
	}
	containerSerialNumber := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(csn + fmt.Sprintf("%v", cd)))

	var bs []rune
	bs = append(bs, []rune(dataIdentifier)...)
	bs = append(bs, ownerCode...)
	bs = append(bs, equipmentIdentifier...)
	bs = append(bs, containerSerialNumber...)

	length := len(dataIdentifier) + len(ownerCode) + len(equipmentIdentifier) + len(containerSerialNumber)
	remainder := length % 16
	var padding []rune
	if remainder != 0 {
		padRuneSlice := binutil.ParseDecimalStringToBinRuneSlice("32") // pad string "100000"
		for i := 0; i < 16-remainder; i++ {
			padding = append(padding, padRuneSlice[i%6])
		}
		bs = append(bs, padding...)
		length += 16 - remainder
	}

	p, err := binutil.ParseBinRuneSliceToUint8Slice(bs)
	if err != nil {
		fmt.Println("Something went wrong!")
		fmt.Println(err)
		os.Exit(1)
	}

	var iso17363 = []interface{}{p}

	return binutil.Pack(iso17363), length
}

// MakeRuneSliceOfISO17365 generates a random 17367 code
func MakeRuneSliceOfISO17365(di string, iac string, cin string, sn string) ([]byte, int) {
	dataIdentifier := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(di))
	issuingAgencyCode := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(iac))
	companyIdentification := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(cin))
	if sn == "" {
		sn = binutil.GenerateNLengthHexString(18)
	}
	serialNumber := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(sn))

	var bs []rune
	bs = append(bs, []rune(dataIdentifier)...)
	bs = append(bs, issuingAgencyCode...)
	bs = append(bs, companyIdentification...)
	bs = append(bs, serialNumber...)

	length := len(dataIdentifier) + len(issuingAgencyCode) + len(companyIdentification) + len(serialNumber)
	remainder := length % 16
	var padding []rune
	if remainder != 0 {
		padRuneSlice := binutil.ParseDecimalStringToBinRuneSlice("32") // pad string "100000"
		for i := 0; i < 16-remainder; i++ {
			padding = append(padding, padRuneSlice[i%6])
		}
		bs = append(bs, padding...)
		length += 16 - remainder
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
