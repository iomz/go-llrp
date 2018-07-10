// Copyright (c) 2018 Iori Mizutani
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package llrp

import (
	"encoding/hex"
	"strconv"
)

// TagRecord stors the Tags contents in string with json tags
type TagRecord struct {
	PCBits string `json:"PCBits"`
	EPC    string `json:"EPC"`
}

func NewTagRecord(tag Tag) *TagRecord {
	return &TagRecord{
		PCBits: strconv.FormatUint(uint64(tag.PCBits), 16),
		EPC:    hex.EncodeToString(tag.EPC),
	}
}
