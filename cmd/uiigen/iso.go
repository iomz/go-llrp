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
func MakeRuneSliceOfISO17363(afi string, oc string, ei string, csn string) ([]byte, int) {
	applicationFamilyIdentifier, _ := binutil.ParseHexStringToBinString(afi)
	ownerCode := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(oc))
	equipmentIdentifier := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(ei))
	fmt.Println(oc + ei + csn)
	cd, err := GetISO6346CD(oc + ei + csn)
	if err != nil {
		panic(err)
	}

	containerSerialNumber := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(csn + fmt.Sprintf("%v", cd)))

	var bs []rune
	bs = append(bs, []rune(applicationFamilyIdentifier)...)
	bs = append(bs, ownerCode...)
	bs = append(bs, equipmentIdentifier...)
	bs = append(bs, containerSerialNumber...)

	length := len(applicationFamilyIdentifier) + len(ownerCode) + len(equipmentIdentifier) + len(containerSerialNumber)
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

	length := len(applicationFamilyIdentifier) + len(dataIdentifier) + len(issuingAgencyCode) + len(companyIdentification) + len(serialNumber)
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
