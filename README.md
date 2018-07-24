go-llrp
==========


[![Build Status](https://travis-ci.org/iomz/go-llrp.svg?branch=master)](https://travis-ci.org/iomz/go-llrp)
[![Coverage Status](https://coveralls.io/repos/iomz/go-llrp/badge.svg?branch=master)](https://coveralls.io/github/iomz/go-llrp?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/iomz/go-llrp)](https://goreportcard.com/report/github.com/iomz/go-llrp)
[![GoDoc](https://godoc.org/github.com/iomz/go-llrp?status.svg)](http://godoc.org/github.com/iomz/go-llrp)

Description
-----------

Tiny llrp library for simple tag data event streaming

Installation
------------

This package can be installed with the go get command:

    go get github.com/iomz/go-llrp

Documentation
-------------

API documentation can be found [here](http://godoc.org/github.com/iomz/go-llrp)
and an example application is [here](https://github.com/iomz/gologir).
Further examples are being prepared soon.

Benchmark
---------

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

Author
------

Iori Mizutani (iomz)

License
-------

```
The MIT License (MIT)
Copyright © 2016 Iori MIZUTANI <iori.mizutani@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
