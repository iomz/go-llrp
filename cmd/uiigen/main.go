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
	epcCompanyPrefix            = epc.Flag("companyPrefix", "Company Prefix for EPC.").Default("").String()
	epcFilter                   = epc.Flag("filter", "Filter Value for EPC.").Default("").String()
	epcItemReference            = epc.Flag("itemReference", "Item Reference Value for EPC.").Default("").String()
	epcExtension                = epc.Flag("extension", "Extension value for EPC.").Default("").String()
	epcSerial                   = epc.Flag("serial", "Serial value for EPC.").Default("").String()
	epcIndivisualAssetReference = epc.Flag("indivisualAssetReference", "Indivisual Asset Reference value for EPC.").Default("").String()
	epcAssetType                = epc.Flag("assetType", "Asset Type for EPC.").Short('y').Default("").String()

	// kingpin generate ISO UII mode
	iso = app.Command("iso", "Generate an ISO UII.")

	// ISO scheme
	isoScheme                      = iso.Flag("scheme", "Scheme for ISO UII.").Default("17365").String()
	isoContainerSerialNumber       = iso.Flag("containerSerialNumber", "A six digit serial number (CSN).").Default("305438").String()
	isoCompanyIdentification       = iso.Flag("companyIdentification", "Company Identification for ISO UII.").Default("043325711").String()
	isoDataIdeintifier             = iso.Flag("dataIdentifier", "Data Identifier for ISO UII.").Default("25S").String()
	isoEquipmentCategoryIdentifier = iso.Flag("equipmentIdentifier", "A one letter equipment category identifier (EI).").Default("U").String()
	isoIssuingAgencyCode           = iso.Flag("issuingAgencyCode", "Issuing Agency Code for ISO UII.").Default("UN").String()
	isoOwnerCode                   = iso.Flag("owenerCode", "A three letter container owner code (OC) assigned in cooperation with the Bureau International des Containers et du Transport Intermodal(BIC).").Default("CSQ").String()
	isoSerialNumber                = iso.Flag("serialNumber", "Serial Number for ISO UII.").Default("MH8031200000000001").String()
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
	switch strings.ToUpper(*epcScheme) {
	case "SGTIN-96":
		uii, _ = MakeRuneSliceOfSGTIN96(*epcCompanyPrefix, *epcFilter, *epcItemReference, *epcSerial)
	case "SSCC-96":
		uii, _ = MakeRuneSliceOfSSCC96(*epcCompanyPrefix, *epcFilter, *epcExtension)
	case "GRAI-96":
		uii, _ = MakeRuneSliceOfGRAI96(*epcCompanyPrefix, *epcFilter, *epcAssetType, *epcSerial)
	case "GIAI-96":
		uii, _ = MakeRuneSliceOfGIAI96(*epcCompanyPrefix, *epcFilter, *epcIndivisualAssetReference)
	}

	// TODO: update pc when length changed (for non-96-bit codes)
	pc := binutil.Pack([]interface{}{
		uint8(48), // L4-0=11000(6words=96bits), UMI=0, XI=0
		uint8(0),  // RFU=0
	})

	length := uint16(18)
	epclen := uint16(96)

	uiibs, _ := binutil.ParseHexStringToBinString(hex.EncodeToString(uii))

	return hex.EncodeToString(pc) + "," +
		strconv.FormatUint(uint64(length), 10) + "," +
		strconv.FormatUint(uint64(epclen), 10) + "," +
		hex.EncodeToString(uii) + "\n" +
		uiibs
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
	case "17365":
		afi := "A2" // 0xA2 ISO 17365 transport uit
		uii, length = MakeRuneSliceOfISO17365(afi, *isoDataIdeintifier, *isoIssuingAgencyCode, *isoCompanyIdentification, *isoSerialNumber)
		pc = MakeISOPC(length, afi)
	case "17363":
		afi := "A9" // 0xA9 ISO 17363 freight containers
		uii, length = MakeRuneSliceOfISO17363(afi, *isoOwnerCode, *isoEquipmentCategoryIdentifier, *isoContainerSerialNumber)
		pc = MakeISOPC(length, afi)
	}

	uiibs, _ := binutil.ParseHexStringToBinString(hex.EncodeToString(uii))

	return hex.EncodeToString(pc) + "," +
		strconv.FormatUint(uint64(length/16), 10) + "," +
		strconv.FormatUint(uint64(length), 10) + "," +
		hex.EncodeToString(uii) + "\n" +
		uiibs
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
