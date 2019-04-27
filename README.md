# go-llrp

[![GoDoc](https://godoc.org/github.com/iomz/go-llrp?status.svg)](http://godoc.org/github.com/iomz/go-llrp)
[![Travis Build Status](https://travis-ci.org/iomz/go-llrp.svg?branch=master)](https://travis-ci.org/iomz/go-llrp)
[![Coverage Status](https://coveralls.io/repos/iomz/go-llrp/badge.svg?branch=master)](https://coveralls.io/github/iomz/go-llrp?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/iomz/go-llrp)](https://goreportcard.com/report/github.com/iomz/go-llrp)
[![GitHub](https://img.shields.io/github/license/iomz/go-llrp.svg)](https://github.com/iomz/go-llrp/blob/master/LICENSE)

The go-llrp package is a tiny library for simple LLRP message and paramter composition.
See [golemu](https://github.com/iomz/golemu) for an example use of this package.

## Installation

Install [dep](https://github.com/golang/dep) first and do the following:

    go get github.com/iomz/go-llrp && cd $GOPATH/src/github.com/iomz/go-llrp && dep ensure

## Benchmark

```bash
% go test -run=XXX -bench=BenchmarkUnmarshal -timeout=0 -v -benchmem
goos: linux
goarch: amd64
pkg: github.com/iomz/go-llrp
BenchmarkUnmarshal100Tags-32              100000             24435 ns/op            7344 B/op        114 allocs/op
BenchmarkUnmarshal200Tags-32               30000             45400 ns/op           14696 B/op        223 allocs/op
BenchmarkUnmarshal300Tags-32               20000             66518 ns/op           22616 B/op        335 allocs/op
BenchmarkUnmarshal400Tags-32               20000             91934 ns/op           29904 B/op        446 allocs/op
BenchmarkUnmarshal500Tags-32               10000            109674 ns/op           37256 B/op        555 allocs/op
BenchmarkUnmarshal600Tags-32               10000            135384 ns/op           45368 B/op        669 allocs/op
BenchmarkUnmarshal700Tags-32               10000            156957 ns/op           52976 B/op        779 allocs/op
BenchmarkUnmarshal800Tags-32               10000            179061 ns/op           59816 B/op        887 allocs/op
BenchmarkUnmarshal900Tags-32               10000            199013 ns/op           67928 B/op       1001 allocs/op
BenchmarkUnmarshal1000Tags-32              10000            227144 ns/op           75536 B/op       1111 allocs/op
PASS
ok      github.com/iomz/go-llrp 82016.381s
go test -run=XXX -bench=BenchmarkUnmarshal -timeout=0 -v -benchmem  158855.53s user 671.25s system 194% cpu 22:46:57.00 total
```

## License

Licensed under the MIT license. Copyright (c) 2016-2019 Iori Mizutani.
