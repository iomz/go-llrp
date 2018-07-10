// Copyright (c) 2018 Iori Mizutani
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package llrp

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"io"
	"os"
)

// Tags holds a slice of pointers to Tag
type Tags []*Tag

// BuildTagReportDataStack takes []*Tag and PDU value to build a new trds
func (tags Tags) BuildTagReportDataStack(pdu int) TagReportDataStack {
	var param []byte
	var trd *TagReportData
	var trds TagReportDataStack
	si := 0 // stack count

	// Iterate through tags and divide them into TRD stacks
	for _, tag := range tags {
		// When exceeds maxTag per TRD, append another TRD in the stack
		// or maximum PDU=int(^uint(0)>>1)
		param = NewTagReportDataParam(tag)
		if len(trds) != 0 &&
			(10+len(trds[si].Data)+4+len(param) >= pdu && int(^uint(0)>>1) > pdu) {
			trd = &TagReportData{Data: param, TagCount: 1}
			trds = append(trds, trd)
			si++
		} else {
			if len(trds) == 0 {
				// First TRD
				trd = &TagReportData{Data: param, TagCount: 1}
				trds = []*TagReportData{trd}
			} else {
				// Append TRD to an existing TRD
				trds[si].Data = append(trds[si].Data, param...)
				trds[si].TagCount++
			}
		}
	}
	return trds
}

// GetIndexOf finds the index in []*Tag
func (tags Tags) GetIndexOf(t *Tag) int {
	index := 0
	for _, tag := range tags {
		if tag.IsDuplicate(t) {
			return index
		}
		index++
	}
	return -1
}

// MarshalBinary overwrites the marshaller in gob encoding Tags
func (tags *Tags) MarshalBinary() (_ []byte, err error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	// Size of tags
	enc.Encode(len(*tags))
	for _, tag := range *tags {
		// Tag
		enc.Encode(tag)
	}

	return buf.Bytes(), err
}

// UnmarshalBinary overwrites the unmarshaller in gob decoding Tags
func (tags *Tags) UnmarshalBinary(data []byte) (err error) {
	dec := gob.NewDecoder(bytes.NewReader(data))

	// Size of Tags
	var tagsSize int
	if err = dec.Decode(&tagsSize); err != nil {
		return
	}

	for i := 0; i < tagsSize; i++ {
		tag := Tag{}
		// Tag
		if err = dec.Decode(&tag); err != nil {
			return
		}
		*tags = append(*tags, &tag)
	}

	return
}

/*
func (tags Tags) writeTagsToCSV(output string) {
	file, err := os.Create(output)
	check(err)

	w := csv.NewWriter(file)
	for _, tag := range tags {
		record := []string{strconv.FormatUint(uint64(tag.PCBits), 16), strconv.FormatUint(uint64(tag.Length), 10), strconv.FormatUint(uint64(tag.EPCLengthBits), 10), hex.EncodeToString(tag.EPC)}
		if err := w.Write(record); err != nil {
			logger.Criticalf("Writing record to csv: %v", err.Error())
		}
		w.Flush()
		if err := w.Error(); err != nil {
			logger.Errorf(err.Error())
		}
	}
	file.Close()
}
*/

// LoadTagsFromCSV reads Tag data from the CSV strings and returns a slice of Tag struct pointers
func LoadTagsFromCSV(inputFile string) Tags {
	// Check inputFile
	fp, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// Read CSV and store in []*Tag
	var tags Tags
	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if len(record) == 2 {
			tagRecord := &TagRecord{record[0], record[1]} // PCbits, EPC
			// Construct a tag read data
			tag, err := NewTag(tagRecord)
			if err != nil {
				continue
			}
			tags = append(tags, tag)
		}
	}

	return tags
}
