// Copyright (c) 2018 Iori Mizutani
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package llrp

import (
	"bytes"
	"encoding/gob"
	"strconv"

	"github.com/iomz/go-llrp/binutil"
)

// Tag holds a single virtual tag content
type Tag struct {
	PCBits uint16
	EPC    []byte
}

// MarshalBinary overwrites the marshaller in gob encoding *Tag
func (tag *Tag) MarshalBinary() (_ []byte, err error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(tag.PCBits)
	enc.Encode(tag.EPC)
	return buf.Bytes(), err
}

// UnmarshalBinary overwrites the unmarshaller in gob decoding *Tag
func (tag *Tag) UnmarshalBinary(data []byte) (err error) {
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err = dec.Decode(&tag.PCBits); err != nil {
		return
	}
	if err = dec.Decode(&tag.EPC); err != nil {
		return
	}
	return
}

// IsEqual to another Tag by taking one as its argument
// return true if they are the same
func (tag *Tag) IsEqual(tt *Tag) bool {
	if tag.PCBits == tt.PCBits && bytes.Equal(tag.EPC, tt.EPC) {
		return true
	}
	return false
}

// IsDuplicate to test another Tag by comparing only EPC
// return true if the EPCs are the same
func (tag *Tag) IsDuplicate(tt *Tag) bool {
	if bytes.Equal(tag.EPC, tt.EPC) {
		return true
	}
	return false
}

// NewTag onstructs a Tag struct from a TagRecord
func NewTag(tagRecord *TagRecord) (*Tag, error) {
	// PCbits
	pc64, err := strconv.ParseUint(tagRecord.PCBits, 16, 16)
	if err != nil {
		return &Tag{}, err
	}
	pc := uint16(pc64)

	// EPC
	epc, err := makeByteID(tagRecord.EPC)
	if err != nil {
		return &Tag{}, err
	}

	return &Tag{pc, epc}, nil
}

func makeByteID(s string) ([]byte, error) {
	id, err := binutil.ParseBinRuneSliceToUint8Slice([]rune(s))
	return binutil.Pack([]interface{}{id}), err
}
