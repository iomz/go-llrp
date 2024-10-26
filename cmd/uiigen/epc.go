// This file contains EPC binary encoding schemes
package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/iomz/go-llrp/binutil"
)

// PartitionTableKey is used for PartitionTables
type PartitionTableKey int

// PartitionTable is used to get the related values for each coding scheme
type PartitionTable map[int]map[PartitionTableKey]int

// Key values for PartitionTables
const (
	PValue PartitionTableKey = iota
	CPBits
	IRBits
	IRDigits
	EBits
	EDigits
	ATBits
	ATDigits
	IARBits
	IARDigits
)

// GIAI96PartitionTable is PT for GIAI
var GIAI96PartitionTable = PartitionTable{
	12: {PValue: 0, CPBits: 40, IARBits: 42, IARDigits: 13},
	11: {PValue: 1, CPBits: 37, IARBits: 45, IARDigits: 14},
	10: {PValue: 2, CPBits: 34, IARBits: 48, IARDigits: 15},
	9:  {PValue: 3, CPBits: 30, IARBits: 52, IARDigits: 16},
	8:  {PValue: 4, CPBits: 27, IARBits: 55, IARDigits: 17},
	7:  {PValue: 5, CPBits: 24, IARBits: 58, IARDigits: 18},
	6:  {PValue: 6, CPBits: 20, IARBits: 62, IARDigits: 19},
}

// GRAI96PartitionTable is PT for GRAI
var GRAI96PartitionTable = PartitionTable{
	12: {PValue: 0, CPBits: 40, ATBits: 4, ATDigits: 0},
	11: {PValue: 1, CPBits: 37, ATBits: 7, ATDigits: 1},
	10: {PValue: 2, CPBits: 34, ATBits: 10, ATDigits: 2},
	9:  {PValue: 3, CPBits: 30, ATBits: 14, ATDigits: 3},
	8:  {PValue: 4, CPBits: 27, ATBits: 17, ATDigits: 4},
	7:  {PValue: 5, CPBits: 24, ATBits: 20, ATDigits: 5},
	6:  {PValue: 6, CPBits: 20, ATBits: 24, ATDigits: 6},
}

// SGTIN96PartitionTable is PT for SGTIN
var SGTIN96PartitionTable = PartitionTable{
	12: {PValue: 0, CPBits: 40, IRBits: 4, IRDigits: 1},
	11: {PValue: 1, CPBits: 37, IRBits: 7, IRDigits: 2},
	10: {PValue: 2, CPBits: 34, IRBits: 10, IRDigits: 3},
	9:  {PValue: 3, CPBits: 30, IRBits: 14, IRDigits: 4},
	8:  {PValue: 4, CPBits: 27, IRBits: 17, IRDigits: 5},
	7:  {PValue: 5, CPBits: 24, IRBits: 20, IRDigits: 6},
	6:  {PValue: 6, CPBits: 20, IRBits: 24, IRDigits: 7},
}

// SSCC96PartitionTable is PT for SSCC
var SSCC96PartitionTable = PartitionTable{
	12: {PValue: 0, CPBits: 40, EBits: 18, EDigits: 5},
	11: {PValue: 1, CPBits: 37, EBits: 21, EDigits: 6},
	10: {PValue: 2, CPBits: 34, EBits: 24, EDigits: 7},
	9:  {PValue: 3, CPBits: 30, EBits: 28, EDigits: 8},
	8:  {PValue: 4, CPBits: 27, EBits: 31, EDigits: 9},
	7:  {PValue: 5, CPBits: 24, EBits: 34, EDigits: 10},
	6:  {PValue: 6, CPBits: 20, EBits: 38, EDigits: 11},
}

// GetAssetType returns Asset Type as rune slice
func GetAssetType(at string, pr map[PartitionTableKey]int) (assetType []rune) {
	if at != "" {
		assetType = binutil.ParseDecimalStringToBinRuneSlice(at)
		if pr[ATBits] > len(assetType) {
			leftPadding := binutil.GenerateNLengthZeroPaddingRuneSlice(pr[ATBits] - len(assetType))
			assetType = append(leftPadding, assetType...)
		}
	} else {
		assetType, _ = binutil.GenerateNLengthRandomBinRuneSlice(pr[ATBits], uint(math.Pow(float64(10), float64(pr[ATDigits]))))
	}
	return
}

// GetCompanyPrefix returns Company Prefix as rune slice
func GetCompanyPrefix(cp string, pt PartitionTable) (companyPrefix []rune) {
	if cp != "" {
		companyPrefix = binutil.ParseDecimalStringToBinRuneSlice(cp)
		if pt[len(cp)][CPBits] > len(companyPrefix) {
			leftPadding := binutil.GenerateNLengthZeroPaddingRuneSlice(pt[len(cp)][CPBits] - len(companyPrefix))
			companyPrefix = append(leftPadding, companyPrefix...)
		}
	}
	return
}

// GetExtension returns Extension digit and Serial Reference as rune slice
func GetExtension(e string, pr map[PartitionTableKey]int) (extension []rune) {
	if e != "" {
		extension = binutil.ParseDecimalStringToBinRuneSlice(e)
		if pr[EBits] > len(extension) {
			leftPadding := binutil.GenerateNLengthZeroPaddingRuneSlice(pr[EBits] - len(extension))
			extension = append(leftPadding, extension...)
		}
	} else {
		extension, _ = binutil.GenerateNLengthRandomBinRuneSlice(pr[EBits], uint(math.Pow(float64(10), float64(pr[EDigits]))))
	}
	return
}

// GetFilter returns filter value as rune slice
func GetFilter(fv string) (filter []rune) {
	if fv != "" {
		n, _ := strconv.ParseInt(fv, 10, 32)
		filter = []rune(fmt.Sprintf("%.3b", n))
	} else {
		filter, _ = binutil.GenerateNLengthRandomBinRuneSlice(3, 7)
	}
	return
}

// GetIndivisualAssetReference returns iar as rune slice
func GetIndivisualAssetReference(iar string, pr map[PartitionTableKey]int) (indivisualAssetReference []rune) {
	if iar != "" {
		indivisualAssetReference = binutil.ParseDecimalStringToBinRuneSlice(iar)
		if pr[IARBits] > len(indivisualAssetReference) {
			leftPadding := binutil.GenerateNLengthZeroPaddingRuneSlice(pr[IARBits] - len(indivisualAssetReference))
			indivisualAssetReference = append(leftPadding, indivisualAssetReference...)
		}
	} else {
		indivisualAssetReference, _ = binutil.GenerateNLengthRandomBinRuneSlice(pr[IARBits], uint(math.Pow(float64(10), float64(pr[IARDigits]))))
	}
	return
}

// GetItemReference converts ItemReference value to rune slice
func GetItemReference(ir string, pr map[PartitionTableKey]int) (itemReference []rune) {
	if ir != "" {
		itemReference = binutil.ParseDecimalStringToBinRuneSlice(ir)
		// If the itemReference is short, pad zeroes to the left
		if pr[IRBits] > len(itemReference) {
			leftPadding := binutil.GenerateNLengthZeroPaddingRuneSlice(pr[IRBits] - len(itemReference))
			itemReference = append(leftPadding, itemReference...)
		}
	} else {
		itemReference, _ = binutil.GenerateNLengthRandomBinRuneSlice(pr[IRBits], uint(math.Pow(float64(10), float64(pr[IRDigits]))))
	}
	return
}

// GetSerial converts serial to rune slice
func GetSerial(s string, serialLength int) (serial []rune) {
	if s != "" {
		serial = binutil.ParseDecimalStringToBinRuneSlice(s)
		if serialLength > len(serial) {
			leftPadding := binutil.GenerateNLengthZeroPaddingRuneSlice(serialLength - len(serial))
			serial = append(leftPadding, serial...)
		}
	} else {
		serial, _ = binutil.GenerateNLengthRandomBinRuneSlice(serialLength, uint(math.Pow(float64(2), float64(serialLength))))
	}
	return serial
}

// MakeGIAI96 generates GIAI-96
func MakeGIAI96(pf bool, fv string, cp string, iar string) ([]byte, string, string, error) {
	filter := GetFilter(fv)
	if fv == "" {
		fv = strconv.Itoa(binutil.ParseBinRuneSliceToInt(filter))
	}

	// CP
	if cp == "" {
		if pf {
			return []byte{}, "00110100" + string(filter), "GIAI-96_" + fv, nil
		}
		return []byte{}, "", "", errors.New("companyPrefix is empty")
	}
	companyPrefix := GetCompanyPrefix(cp, GIAI96PartitionTable)
	partition := []rune(fmt.Sprintf("%.3b", GIAI96PartitionTable[len(cp)][PValue]))

	// IAR
	if iar == "" {
		if pf {
			return []byte{}, "00110100" + string(filter) + string(partition) + string(companyPrefix), "GIAI-96_" + fv + "_" + strconv.Itoa(GIAI96PartitionTable[len(cp)][PValue]) + "_" + cp, nil
		}
	}
	indivisualAssetReference := GetIndivisualAssetReference(iar, GIAI96PartitionTable[len(cp)])

	// Exact match
	if pf {
		return []byte{}, "00110100" + string(filter) + string(partition) + string(companyPrefix) + string(indivisualAssetReference), "GIAI-96_" + fv + "_" + strconv.Itoa(GIAI96PartitionTable[len(cp)][PValue]) + "_" + cp + "_" + iar, nil
	}

	bs := append(filter, partition...)
	bs = append(bs, companyPrefix...)
	bs = append(bs, indivisualAssetReference...)

	if len(bs) != 88 {
		return []byte{}, "", "", fmt.Errorf("len(bs): %d", len(bs))
	}

	p, err := binutil.ParseBinRuneSliceToUint8Slice(bs)
	if err != nil {
		return []byte{}, "", "", err
	}

	var giai96 = []interface{}{
		uint8(52), // GIAI-96 Header 0011 0100
		p[0],      // 8 bits -> 16 bits
		p[1],      // 8 bits -> 24 bits
		p[2],      // 8 bits -> 32 bits
		p[3],      // 8 bits -> 40 bits
		p[4],      // 8 bits -> 48 bits
		p[5],      // 8 bits -> 56 bits
		p[6],      // 8 bits -> 64 bits
		p[7],      // 8 bits -> 72 bits
		p[8],      // 8 bits -> 80 bits
		p[9],      // 8 bits -> 88 bits
		p[10],     // 8 bits -> 96 bits
	}

	return binutil.Pack(giai96), "", "", nil
}

// MakeGRAI96 generates GRAI-96
func MakeGRAI96(pf bool, fv string, cp string, at string, ser string) ([]byte, string, string, error) {
	filter := GetFilter(fv)

	//CP
	if cp == "" {
		if pf {
			return []byte{}, "00110011" + string(filter), "GRAI-96_" + fv, nil
		}
		return []byte{}, "", "", errors.New("companyPrefix is empty")
	}
	companyPrefix := GetCompanyPrefix(cp, GRAI96PartitionTable)
	partition := []rune(fmt.Sprintf("%.3b", GRAI96PartitionTable[len(cp)][PValue]))

	//AT
	if at == "" {
		if pf {
			return []byte{}, "00110011" + string(filter) + string(partition) + string(companyPrefix), "GRAI-96_" + fv + "_" + strconv.Itoa(GRAI96PartitionTable[len(cp)][PValue]) + "_" + cp, nil
		}
	}
	assetType := GetAssetType(at, GRAI96PartitionTable[len(cp)])

	// SER
	if ser == "" {
		if pf {
			return []byte{}, "00110011" + string(filter) + string(partition) + string(companyPrefix) + string(assetType), "GRAI-96_" + fv + "_" + strconv.Itoa(GRAI96PartitionTable[len(cp)][PValue]) + "_" + cp + "_" + at, nil
		}
	}
	serial := GetSerial(ser, 38)

	// Exact match
	if pf {
		return []byte{}, "00110011" + string(filter) + string(partition) + string(companyPrefix) + string(assetType) + string(serial), "GRAI-96_" + fv + "_" + strconv.Itoa(GRAI96PartitionTable[len(cp)][PValue]) + "_" + cp + "_" + at + "_" + ser, nil
	}

	bs := append(filter, partition...)
	bs = append(bs, companyPrefix...)
	bs = append(bs, assetType...)
	bs = append(bs, serial...)

	if len(bs) != 88 {
		return []byte{}, "", "", fmt.Errorf("len(bs): %d", len(bs))
	}

	p, err := binutil.ParseBinRuneSliceToUint8Slice(bs)
	if err != nil {
		return []byte{}, "", "", err
	}

	var grai96 = []interface{}{
		uint8(51), // GRAI-96 Header 0011 0011
		p[0],      // 8 bits -> 16 bits
		p[1],      // 8 bits -> 24 bits
		p[2],      // 8 bits -> 32 bits
		p[3],      // 8 bits -> 40 bits
		p[4],      // 8 bits -> 48 bits
		p[5],      // 8 bits -> 56 bits
		p[6],      // 8 bits -> 64 bits
		p[7],      // 8 bits -> 72 bits
		p[8],      // 8 bits -> 80 bits
		p[9],      // 8 bits -> 88 bits
		p[10],     // 8 bits -> 96 bits
	}

	return binutil.Pack(grai96), "", "", nil
}

// MakeSGTIN96 generates SGTIN-96
func MakeSGTIN96(pf bool, fv string, cp string, ir string, ser string) ([]byte, string, string, error) {
	filter := GetFilter(fv)
	// CP
	if cp == "" {
		if pf {
			return []byte{}, "00110000" + string(filter), "SGTIN-96_" + fv, nil
		}
		return []byte{}, "", "", errors.New("companyPrefix is empty")
	}
	companyPrefix := GetCompanyPrefix(cp, SGTIN96PartitionTable)
	partition := []rune(fmt.Sprintf("%.3b", SGTIN96PartitionTable[len(cp)][PValue]))

	// IR
	if ir == "" {
		if pf {
			return []byte{}, "00110000" + string(filter) + string(partition) + string(companyPrefix), "SGTIN-96_" + fv + "_" + strconv.Itoa(SGTIN96PartitionTable[len(cp)][PValue]) + "_" + cp, nil
		}
	}
	itemReference := GetItemReference(ir, SGTIN96PartitionTable[len(cp)])

	// SER
	if ser == "" {
		if pf {
			return []byte{}, "00110000" + string(filter) + string(partition) + string(companyPrefix) + string(itemReference), "SGTIN-96_" + fv + "_" + strconv.Itoa(SGTIN96PartitionTable[len(cp)][PValue]) + "_" + cp + "_" + ir, nil
		}
	}
	serial := GetSerial(ser, 38)

	// Exact match
	if pf {
		return []byte{}, "00110000" + string(filter) + string(partition) + string(companyPrefix) + string(itemReference) + string(serial), "SGTIN-96_" + fv + "_" + strconv.Itoa(SGTIN96PartitionTable[len(cp)][PValue]) + "_" + cp + "_" + ir + "_" + ser, nil
	}

	bs := append(filter, partition...)
	bs = append(bs, companyPrefix...)
	bs = append(bs, itemReference...)
	bs = append(bs, serial...)

	if len(bs) != 88 {
		return []byte{}, "", "", fmt.Errorf("len(bs): %d", len(bs))
	}

	p, err := binutil.ParseBinRuneSliceToUint8Slice(bs)
	if err != nil {
		return []byte{}, "", "", err
	}

	var sgtin96 = []interface{}{
		uint8(48), // SGTIN-96 Header 0011 0000
		p[0],      // 8 bits -> 16 bits
		p[1],      // 8 bits -> 24 bits
		p[2],      // 8 bits -> 32 bits
		p[3],      // 8 bits -> 40 bits
		p[4],      // 8 bits -> 48 bits
		p[5],      // 8 bits -> 56 bits
		p[6],      // 8 bits -> 64 bits
		p[7],      // 8 bits -> 72 bits
		p[8],      // 8 bits -> 80 bits
		p[9],      // 8 bits -> 88 bits
		p[10],     // 8 bits -> 96 bits
	}

	return binutil.Pack(sgtin96), "", "", nil
}

// MakeSSCC96 generates SSCC-96
func MakeSSCC96(pf bool, fv string, cp string, ext string) ([]byte, string, string, error) {
	filter := GetFilter(fv)
	// CP
	if cp == "" {
		if pf {
			return []byte{}, "00110001" + string(filter), "SSCC-96_" + fv, nil
		}
		return []byte{}, "", "", errors.New("companyPrefix is empty")
	}
	companyPrefix := GetCompanyPrefix(cp, SSCC96PartitionTable)
	partition := []rune(fmt.Sprintf("%.3b", SSCC96PartitionTable[len(cp)][PValue]))

	// EXT
	if ext == "" {
		if pf {
			return []byte{}, "00110001" + string(filter) + string(partition) + string(companyPrefix), "SSCC-96_" + fv + "_" + strconv.Itoa(SSCC96PartitionTable[len(cp)][PValue]) + "_" + cp, nil
		}
	}
	extension := GetExtension(ext, SSCC96PartitionTable[len(cp)])

	// Exact match (ignore rsvd
	if pf {
		return []byte{}, "00110001" + string(filter) + string(partition) + string(companyPrefix) + string(extension), "SSCC-96_" + fv + "_" + strconv.Itoa(SSCC96PartitionTable[len(cp)][PValue]) + "_" + cp + "_" + ext, nil
	}

	// 24 '0's
	reserved := binutil.GenerateNLengthZeroPaddingRuneSlice(24)

	bs := append(filter, partition...)
	bs = append(bs, companyPrefix...)
	bs = append(bs, extension...)
	bs = append(bs, reserved...)

	if len(bs) != 88 {
		return []byte{}, "", "", fmt.Errorf("len(bs): %d", len(bs))
	}

	p, err := binutil.ParseBinRuneSliceToUint8Slice(bs)
	if err != nil {
		return []byte{}, "", "", err
	}

	var sscc96 = []interface{}{
		uint8(49), // SSCC-96 Header 0011 0001
		p[0],      // 8 bits -> 16 bits
		p[1],      // 8 bits -> 24 bits
		p[2],      // 8 bits -> 32 bits
		p[3],      // 8 bits -> 40 bits
		p[4],      // 8 bits -> 48 bits
		p[5],      // 8 bits -> 56 bits
		p[6],      // 8 bits -> 64 bits
		p[7],      // 8 bits -> 72 bits
		p[8],      // 8 bits -> 80 bits
		p[9],      // 8 bits -> 88 bits
		p[10],     // 8 bits -> 96 bits
	}

	return binutil.Pack(sscc96), "", "", nil
}
