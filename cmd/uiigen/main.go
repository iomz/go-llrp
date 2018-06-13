// A tool to generate arbitrary UII (aka EPC)
package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iomz/go-llrp/binutil"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// kingpin app
	app          = kingpin.New("uiigen", "A tool to generate an arbitrary UII (aka EPC).")
	prefixFilter = app.Flag("pf", "Print a prefix filter for the given parameter").Default("false").Bool()
	modeHex      = app.Flag("hex", "Print the ID in Hex.").Default("false").Bool()
	modeDec      = app.Flag("dec", "Print the ID in Dec.").Default("false").Bool()

	// kingpin generate EPC mode
	epc = app.Command("epc", "Generate an EPC.")
	// EPC scheme
	epcScheme                   = epc.Flag("cs", "EPC coding scheme. ex.) sgtin-96").String()
	epcFilter                   = epc.Flag("filter", "Filter Value for EPC.").Default("3").String()
	epcCompanyPrefix            = epc.Flag("cp", "[GIAI,GRAI,SGTIN,SSCC] Company Prefix for EPC. ex.) 0614141").String()
	epcItemReference            = epc.Flag("ir", "[SGTIN] Item Reference Value for EPC.").String()
	epcExtension                = epc.Flag("ext", "[SSCC] Extension value for EPC.").String()
	epcSerial                   = epc.Flag("ser", "[GRAI,SGTIN] Serial value for EPC.").String()
	epcIndivisualAssetReference = epc.Flag("iar", "[GIAI] Indivisual Asset Reference value for EPC.").String()
	epcAssetType                = epc.Flag("at", "[GRAI] Asset Type for EPC.").String()

	// kingpin generate ISO UII mode
	iso = app.Command("iso", "Generate an ISO UII.")
	// ISO scheme
	iso17363                       = iso.Command("17363", "ISO 17363 coding scheme.")
	isoOwnerCode                   = iso17363.Flag("oc", "A three letter container owner code (OC) assigned in cooperation with the Bureau International des Containers et du Transport Intermodal(BIC). ex.) CSQ").String()
	isoEquipmentCategoryIdentifier = iso17363.Flag("ei", "A one letter equipment category identifier (EI) in ISO 6346 = U, J or Z.").Default("U").String()
	isoContainerSerialNumber       = iso17363.Flag("csn", "A six digit serial number (CSN). ex.) 305438").String()
	iso17365                       = iso.Command("17365", "ISO 17365 coding scheme.")
	isoDataIdeintifier             = iso17365.Flag("di", "Data Identifier in ISO/IEC 15418. ex.) 25S").Default("25S").String()
	isoIssuingAgencyCode           = iso17365.Flag("iac", "1-3 Alphabet letters for Issuing Agency Code in ISO 15459. ex.) UN or OD").String()
	isoCompanyIdentification       = iso17365.Flag("cin", "Company Identification. ex.) 043325711 or CIN1").String()
	isoSerialNumber                = iso17365.Flag("sn", "Serial Number for ISO UII. ex.) MH8031200000000001 or 0000000RTIA1B2C3DOSN12345").String()
)

// CheckIfStringInSlice checks if string exists in a string slice
// TODO: fix the way it is, it should be smarter
func CheckIfStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// MakeEPC generates EPC in binary string and PC in hex string or prefix and elements
func MakeEPC(pf bool, cs string, fv string, cp string, ir string, ext string, ser string, iar string, at string) (string, string) {
	var uii []byte
	var f string
	var elem string

	switch strings.ToUpper(cs) {
	case "GIAI-96":
		uii, f, elem, _ = MakeGIAI96(pf, fv, cp, iar)
	case "GRAI-96":
		uii, f, elem, _ = MakeGRAI96(pf, fv, cp, at, ser)
	case "SGTIN-96":
		uii, f, elem, _ = MakeSGTIN96(pf, fv, cp, ir, ser)
	case "SSCC-96":
		uii, f, elem, _ = MakeSSCC96(pf, fv, cp, ext)
	}

	// If only prefix flag is on, return prefix as epc
	if pf {
		return f, elem
	}

	// TODO: update pc when length changed (for non-96-bit codes)
	pc := binutil.Pack([]interface{}{
		uint8(48), // L4-0=11000(6words=96bits), UMI=0, XI=0
		uint8(0),  // RFU=0
	})

	uiibs, _ := binutil.ParseHexStringToBinString(hex.EncodeToString(uii))

	return uiibs, hex.EncodeToString(pc)
	/*
		length := uint16(18)
		epclen := uint16(96)
	*/
}

// MakeISO returns ISO UII in binary string and PCbits in hex string or prefix and elements
func MakeISO(pf bool, std string, oc string, ei string, csn string, di string, iac string, cin string, sn string) (string, string) {
	var uii []byte
	var pc []byte
	var length int
	var f string
	var elem string

	switch std {
	case "17363":
		afi := "A9" // 0xA9 ISO 17363 freight containers
		uii, length, f, elem, _ = MakeISO17363(pf, oc, ei, csn)
		pc = MakeISOPC(length, afi)
	case "17365":
		afi := "A2" // 0xA2 ISO 17365 transport uit
		uii, length, f, elem, _ = MakeISO17365(pf, di, iac, cin, sn)
		pc = MakeISOPC(length, afi)
	}

	// If only prefix flag is on, return prefix as iso uii
	if pf {
		return f, elem
	}

	uiibs, _ := binutil.ParseHexStringToBinString(hex.EncodeToString(uii))

	return uiibs, hex.EncodeToString(pc)
	/*
		return hex.EncodeToString(pc) + "," +
			strconv.FormatUint(uint64(length/16), 10) + "," +
			strconv.FormatUint(uint64(length), 10) + "," +
			hex.EncodeToString(uii) + "\n" +
			uiibs
	*/
}

// MakeISOPC returns PC bits in []byte
func MakeISOPC(length int, afi string) []byte {
	l := []rune(fmt.Sprintf("%.5b", length/16))
	pc1, err := binutil.ParseBinRuneSliceToUint8Slice(append(l, rune('0'), rune('0'), rune('1'))) // L, UMI, XI, T
	if err != nil {
		panic(err)
	}
	c, _ := strconv.ParseUint(afi, 16, 8)
	return binutil.Pack([]interface{}{
		pc1[0],
		uint8(c), // AFI
	})
}

func main() {
	parse := kingpin.MustParse(app.Parse(os.Args[1:]))

	var bs string
	var opt string
	switch parse {
	case epc.FullCommand():
		bs, opt = MakeEPC(*prefixFilter, *epcScheme, *epcFilter, *epcCompanyPrefix, *epcItemReference, *epcExtension, *epcSerial, *epcIndivisualAssetReference, *epcAssetType)
	case iso17363.FullCommand():
		bs, opt = MakeISO(*prefixFilter, "17363", *isoOwnerCode, *isoEquipmentCategoryIdentifier, *isoContainerSerialNumber, *isoDataIdeintifier, *isoIssuingAgencyCode, *isoCompanyIdentification, *isoSerialNumber)
	case iso17365.FullCommand():
		bs, opt = MakeISO(*prefixFilter, "17365", *isoOwnerCode, *isoEquipmentCategoryIdentifier, *isoContainerSerialNumber, *isoDataIdeintifier, *isoIssuingAgencyCode, *isoCompanyIdentification, *isoSerialNumber)
	}
	if len(bs) != 0 {
		if *modeHex && !*prefixFilter {
			hs, _ := binutil.ParseBinStringToHexString(bs)
			fmt.Println(opt + "," + hs)
		} else if *modeDec {
			ds, _ := binutil.ParseBinStringToDecArrayString(bs)
			pc, _ := binutil.ParseHexStringToDecArrayString(opt)
			fmt.Println("[]byte{" + pc + "},[]byte{" + ds + "}")
		} else {
			fmt.Println(opt + "," + bs)
		}
	}
}
