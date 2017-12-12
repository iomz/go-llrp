// This file contains ISO binary encoding scheme
package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iomz/go-llrp/binutil"
)

// GetISO6346CD returns check digit for container serial number
func GetISO6346CD(cn string) (int, error) {
	if len(cn) != 10 {
		return 0, errors.New("Invalid ISO6346 code provided")
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

// MakeISO17363 generates a random 17363 code
func MakeISO17363(pf bool, oc string, ei string, csn string) ([]byte, int, string, string, error) {
	di := "7B"
	dataIdentifier := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(di))

	// OC
	if oc == "" {
		if pf {
			return []byte{}, 0, string(dataIdentifier), "ISO17363_" + di, nil
		}
		oc = binutil.GenerateNLengthAlphabetString(3)
	}
	ownerCode := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(oc))

	// EI
	if ei == "" {
		if pf {
			return []byte{}, 0, string(dataIdentifier) + string(ownerCode), "ISO17363_" + di + "_" + oc, nil
		}
		ei = "U"
	}
	equipmentIdentifier := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(ei))

	// CSN
	if csn == "" {
		if pf {
			return []byte{}, 0, string(dataIdentifier) + string(ownerCode) + string(equipmentIdentifier), "ISO17363_" + di + "_" + oc + "_" + ei, nil
		}
		csn = binutil.GenerateNLengthDigitString(6)
	} else if 6 > len(csn) {
		leftPadding := binutil.GenerateNLengthZeroPaddingRuneSlice(6 - len(csn))
		csn = string(leftPadding) + csn
	} else if 6 < len(csn) {
		return []byte{}, 0, "", "", errors.New("Invalid csn: " + csn)
	}
	cd, err := GetISO6346CD(oc + ei + csn)
	if err != nil {
		return []byte{}, 0, "", "", err
	}
	containerSerialNumber := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(csn + fmt.Sprintf("%v", cd)))

	// Exact match filter
	if pf {
		return []byte{}, 0, string(dataIdentifier) + string(ownerCode) + string(equipmentIdentifier) + string(containerSerialNumber), "ISO17363_" + di + "_" + oc + "_" + ei + "_" + csn, nil
	}

	bs := append(dataIdentifier, ownerCode...)
	bs = append(bs, equipmentIdentifier...)
	bs = append(bs, containerSerialNumber...)

	var length int
	bs, length = Pad6BitEncodingRuneSlice(bs)

	p, err := binutil.ParseBinRuneSliceToUint8Slice(bs)
	if err != nil {
		return []byte{}, 0, "", "", err
	}

	var iso17363 = []interface{}{p}

	return binutil.Pack(iso17363), length, "", "", nil
}

// MakeISO17365 generates a random 17367 code
func MakeISO17365(pf bool, di string, iac string, cin string, sn string) ([]byte, int, string, string, error) {
	dataIdentifier := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(di))

	// IAC
	if iac == "" {
		if pf {
			return []byte{}, 0, string(dataIdentifier), "ISO17365_" + di, nil
		}
		return []byte{}, 0, "", "", errors.New("IAC not provided")
	}
	issuingAgencyCode := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(iac))

	// CIN
	if cin == "" {
		if pf {
			return []byte{}, 0, string(dataIdentifier) + string(issuingAgencyCode), "ISO17365_" + di + "_" + iac, nil
		}
		return []byte{}, 0, "", "", errors.New("CIN not provided")
	}
	companyIdentification := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(cin))

	// SN
	if sn == "" {
		if pf {
			return []byte{}, 0, string(dataIdentifier) + string(issuingAgencyCode) + string(companyIdentification), "ISO17365_" + di + "_" + iac + "_" + cin, nil
		}
		sn = binutil.GenerateNLengthHexString(18)
	}
	serialNumber := binutil.ParseRuneSliceTo6BinRuneSlice([]rune(sn))

	// Exact match filter
	if pf {
		return []byte{}, 0, string(dataIdentifier) + string(issuingAgencyCode) + string(companyIdentification) + string(serialNumber), "ISO17365_" + di + "_" + iac + "_" + cin + "_" + sn, nil
	}

	bs := append(dataIdentifier, issuingAgencyCode...)
	bs = append(bs, companyIdentification...)
	bs = append(bs, serialNumber...)

	var length int
	bs, length = Pad6BitEncodingRuneSlice(bs)

	p, err := binutil.ParseBinRuneSliceToUint8Slice(bs)
	if err != nil {
		return []byte{}, 0, "", "", err
	}

	var iso17365 = []interface{}{p}

	return binutil.Pack(iso17365), length, "", "", nil
}

// Pad6BitEncodingRuneSlice returns a new length
// and 16-bit (word-length) padded binary string in rune slice
// @ISO15962
func Pad6BitEncodingRuneSlice(bs []rune) ([]rune, int) {
	length := len(bs)
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
	return bs, length
}
