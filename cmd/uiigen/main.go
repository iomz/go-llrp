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
	app = kingpin.New("uiigen", "A tool to generate an arbitrary UII (aka EPC).")

	// kingpin generate EPC mode
	epc = app.Command("epc", "Generate an EPC.")
	// EPC scheme
	epcScheme                   = epc.Flag("type", "EPC UII type.").Default("SGTIN-96").String()
	epcFilter                   = epc.Flag("filter", "Filter Value for EPC.").Default("3").String()
	epcCompanyPrefix            = epc.Flag("companyPrefix", "[GIAI,GRAI,SGTIN,SSCC] Company Prefix for EPC.").Default("0614141").String()
	epcItemReference            = epc.Flag("itemReference", "[SGTIN] Item Reference Value for EPC.").String()
	epcExtension                = epc.Flag("extension", "[SSCC] Extension value for EPC.").String()
	epcSerial                   = epc.Flag("serial", "[GRAI,SGTIN] Serial value for EPC.").String()
	epcIndivisualAssetReference = epc.Flag("indivisualAssetReference", "[GIAI] Indivisual Asset Reference value for EPC.").String()
	epcAssetType                = epc.Flag("assetType", "[GRAI] Asset Type for EPC.").String()

	// kingpin generate ISO UII mode
	iso = app.Command("iso", "Generate an ISO UII.")

	// ISO scheme
	isoScheme                      = iso.Flag("scheme", "Scheme for ISO UII.").Default("17365").String()
	isoOwnerCode                   = iso.Flag("ownerCode", "[17363] A three letter container owner code (OC) assigned in cooperation with the Bureau International des Containers et du Transport Intermodal(BIC). ex.) CSQ").String()
	isoEquipmentCategoryIdentifier = iso.Flag("equipmentIdentifier", "[17363] A one letter equipment category identifier (EI).").Default("U").String()
	isoContainerSerialNumber       = iso.Flag("containerSerialNumber", "[17363] A six digit serial number (CSN). ex.) 305438").String()
	isoDataIdeintifier             = iso.Flag("dataIdentifier", "[17365] Data Identifier for ISO UII.").Default("25S").String()
	isoIssuingAgencyCode           = iso.Flag("issuingAgencyCode", "[17365] Issuing Agency Code for ISO UII.").Default("UN").String()
	isoCompanyIdentification       = iso.Flag("companyIdentification", "[17365] Company Identification for ISO UII.").Default("043325711").String()
	isoSerialNumber                = iso.Flag("serialNumber", "[17365] Serial Number for ISO UII. ex.) MH8031200000000001").String()
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

// MakeEPC generates EPC code
func MakeEPC() string {
	epcs := []string{"SGTIN-96", "SSCC-96", "GRAI-96", "GIAI-96"}

	if !CheckIfStringInSlice(strings.ToUpper(*epcScheme), epcs) {
		os.Exit(1)
	}

	var uii []byte
	var err error
	switch strings.ToUpper(*epcScheme) {
	case "GIAI-96":
		uii, err = MakeRuneSliceOfGIAI96(*epcCompanyPrefix, *epcFilter, *epcIndivisualAssetReference)
	case "GRAI-96":
		uii, err = MakeRuneSliceOfGRAI96(*epcCompanyPrefix, *epcFilter, *epcAssetType, *epcSerial)
	case "SGTIN-96":
		uii, err = MakeRuneSliceOfSGTIN96(*epcCompanyPrefix, *epcFilter, *epcItemReference, *epcSerial)
	case "SSCC-96":
		uii, err = MakeRuneSliceOfSSCC96(*epcCompanyPrefix, *epcFilter, *epcExtension)
	}
	if err != nil {
		panic(err)
	}

	// TODO: update pc when length changed (for non-96-bit codes)
	pc := binutil.Pack([]interface{}{
		uint8(48), // L4-0=11000(6words=96bits), UMI=0, XI=0
		uint8(0),  // RFU=0
	})

	uiibs, _ := binutil.ParseHexStringToBinString(hex.EncodeToString(uii))

	return hex.EncodeToString(pc) + "," + uiibs
	/*
	length := uint16(18)
	epclen := uint16(96)
	return hex.EncodeToString(pc) + "," +
		strconv.FormatUint(uint64(length), 10) + "," +
		strconv.FormatUint(uint64(epclen), 10) + "," +
		hex.EncodeToString(uii) + "\n" +
		uiibs
	*/
}

// MakeISO returns ISO code
func MakeISO() string {
	var uii []byte
	var pc []byte
	var length int

	isos := []string{"17365", "17363"}

	if !CheckIfStringInSlice(*isoScheme, isos) {
		os.Exit(1)
	}

	switch *isoScheme {
	case "17363":
		afi := "A9" // 0xA9 ISO 17363 freight containers
		uii, length = MakeRuneSliceOfISO17363(afi, *isoOwnerCode, *isoEquipmentCategoryIdentifier, *isoContainerSerialNumber)
		pc = MakeISOPC(length, afi)
	case "17365":
		afi := "A2" // 0xA2 ISO 17365 transport uit
		uii, length = MakeRuneSliceOfISO17365(afi, *isoDataIdeintifier, *isoIssuingAgencyCode, *isoCompanyIdentification, *isoSerialNumber)
		pc = MakeISOPC(length, afi)
	}

	uiibs, _ := binutil.ParseHexStringToBinString(hex.EncodeToString(uii))

	return hex.EncodeToString(pc) + "," + uiibs
	/*
	return hex.EncodeToString(pc) + "," +
		strconv.FormatUint(uint64(length/16), 10) + "," +
		strconv.FormatUint(uint64(length), 10) + "," +
		hex.EncodeToString(uii) + "\n" +
		uiibs
	*/
}

// MakePC returns PC bits
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

	switch parse {
	case epc.FullCommand():
		fmt.Println(MakeEPC())
	case iso.FullCommand():
		fmt.Println(MakeISO())
	}
}
