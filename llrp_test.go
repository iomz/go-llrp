// Copyright (c) 2018 Iori Mizutani
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package llrp

import (
	"bytes"
	"math/rand"
	"os"
	"testing"

	"github.com/iomz/go-llrp/binutil"
)

var packtests = []struct {
	in  []interface{}
	out []byte
}{
	{[]interface{}{uint16(349), uint16(11), uint8(0)}, []byte{1, 93, 0, 11, 0}},
	{[]interface{}{uint8(12), uint8(11), uint32(433)}, []byte{12, 11, 0, 0, 1, 177}},
}

func TestPack(t *testing.T) {
	var b []byte
	for _, tt := range packtests {
		b = Pack(tt.in)
		if !bytes.Equal(b, tt.out) {
			t.Errorf("%v => %v, want %v", tt.in, b, tt.out)
		}
	}
}

func TestUnmarshalROAccessReportBody(t *testing.T) {
	largeTagsGOB := os.Getenv("GOPATH") + "/src/github.com/iomz/go-llrp/test/data/1000-tags.gob"
	size := 100
	// load up the tags from the file
	var largeTags Tags
	binutil.Load(largeTagsGOB, &largeTags)

	// cap the tags with the given size
	var limitedTags Tags
	perms := rand.Perm(len(largeTags))
	for n, i := range perms {
		if n < size {
			limitedTags = append(limitedTags, largeTags[i])
		} else {
			break
		}
		if n+1 == len(largeTags) {
			t.Fatal("given tag size is larger than the testdata available")
		}
	}

	// build ROAR message
	pdu := int(1500)
	trds := limitedTags.BuildTagReportDataStack(pdu)
	if len(trds) == 0 {
		t.Fatal("TagReportDataStack generation failed")
	}

	var res []*ReadEvent
	for i, trd := range trds {
		roar := NewROAccessReport(trd.Data, uint32(i))
		res = append(res, UnmarshalROAccessReportBody(roar.data[10:])...)
	}

	if len(res) != size {
		t.Errorf("UnmarshalROAccessReport() = %v", res)
	}
}

/*
func BenchmarkUnmarshalLargeROAR(b *testing.B) {
	largeTagsGOB := os.Getenv("GOPATH") + "/src/github.com/iomz/go-llrp/test/data/million-tags.gob"
	// load up the tags from the file
	var largeTags Tags
	binutil.Load(largeTagsGOB, &largeTags)

	cycle := b.N / len(largeTags)
	remaining := b.N % len(largeTags)

	// cap the tags with the given size
	var limitedTags Tags
	perms := rand.Perm(len(largeTags))
	for n, i := range perms {
		if n < remaining {
			limitedTags = append(limitedTags, largeTags[i])
		} else {
			break
		}
		if n == len(largeTags) {
			b.Skip("given tag size is larger than the testdata available")
		}
	}

	// build ROAR message
	pdu := int(1500)
	trds := largeTags.BuildTagReportDataStack(pdu)
	if len(trds) == 0 {
		b.Fatal("TagReportDataStack generation was failed")
	}
	limitedTRDs := limitedTags.BuildTagReportDataStack(pdu)
	if len(limitedTRDs) == 0 && remaining != 0 {
		b.Logf("len(limitedTags): %v, len(limitedTRDs: %v", len(limitedTags), len(limitedTRDs))
		b.Logf("b.N: %v, cycle: %v, remaining: %v", b.N, cycle, remaining)
		b.Fatal("TagReportDataStack generation failed")
	}

	var res []*ReadEvent
	b.ResetTimer()
	for c := 0; c < cycle; c++ {
		for i, trd := range trds {
			b.StopTimer()
			roar := NewROAccessReport(trd.Data, uint32(i))
			b.StartTimer()
			res = append(res, UnmarshalROAccessReportBody(roar.data[10:])...)
		}
	}

	for i, trd := range limitedTRDs {
		b.StopTimer()
		roar := NewROAccessReport(trd.Data, uint32(i))
		b.StartTimer()
		res = append(res, UnmarshalROAccessReportBody(roar.data[10:])...)
	}
	b.StopTimer()
	if b.N != len(res) {
		b.Fatal("LLRP unmarshaller failed")
	}
}
*/

func benchmarkUnmarshalNROARNTags(nROAR int, nTags int, b *testing.B) {
	largeTagsGOB := os.Getenv("GOPATH") + "/src/github.com/iomz/go-llrp/test/data/million-tags.gob"
	// load up the tags from the file
	var largeTags Tags
	binutil.Load(largeTagsGOB, &largeTags)

	// cap the tags with the given size
	var limitedTags Tags
	perms := rand.Perm(len(largeTags))
	for count, i := range perms {
		if count < nTags {
			limitedTags = append(limitedTags, largeTags[i])
		} else {
			break
		}
		if count == len(largeTags) {
			b.Skip("given tag size is larger than the testdata available")
		}
	}
	// build ROAR message
	trds := limitedTags.BuildTagReportDataStack(1500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, trd := range trds {
			b.StopTimer()
			roar := NewROAccessReport(trd.Data, uint32(i))
			b.StartTimer()
			res := UnmarshalROAccessReportBody(roar.data[10:])
			_ = res
		}
	}
}

func BenchmarkUnmarshal100Tags(b *testing.B)  { benchmarkUnmarshalNROARNTags(100, 100, b) }
func BenchmarkUnmarshal200Tags(b *testing.B)  { benchmarkUnmarshalNROARNTags(200, 100, b) }
func BenchmarkUnmarshal300Tags(b *testing.B)  { benchmarkUnmarshalNROARNTags(300, 100, b) }
func BenchmarkUnmarshal400Tags(b *testing.B)  { benchmarkUnmarshalNROARNTags(400, 100, b) }
func BenchmarkUnmarshal500Tags(b *testing.B)  { benchmarkUnmarshalNROARNTags(500, 100, b) }
func BenchmarkUnmarshal600Tags(b *testing.B)  { benchmarkUnmarshalNROARNTags(600, 100, b) }
func BenchmarkUnmarshal700Tags(b *testing.B)  { benchmarkUnmarshalNROARNTags(700, 100, b) }
func BenchmarkUnmarshal800Tags(b *testing.B)  { benchmarkUnmarshalNROARNTags(800, 100, b) }
func BenchmarkUnmarshal900Tags(b *testing.B)  { benchmarkUnmarshalNROARNTags(900, 100, b) }
func BenchmarkUnmarshal1000Tags(b *testing.B) { benchmarkUnmarshalNROARNTags(1000, 100, b) }
