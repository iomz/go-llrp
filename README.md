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
BenchmarkUnmarshal100ROAR100Tags-32                 1000           1291573 ns/op          684000 B/op      10800 allocs/op
BenchmarkUnmarshal200ROAR100Tags-32                  500           2644347 ns/op         1368000 B/op      21600 allocs/op
BenchmarkUnmarshal300ROAR100Tags-32                  500           4090107 ns/op         2052000 B/op      32400 allocs/op
BenchmarkUnmarshal400ROAR100Tags-32                  300           5539019 ns/op         2736000 B/op      43200 allocs/op
BenchmarkUnmarshal500ROAR100Tags-32                  200           6885189 ns/op         3420000 B/op      54000 allocs/op
BenchmarkUnmarshal600ROAR100Tags-32                  200           8500495 ns/op         4104000 B/op      64800 allocs/op
BenchmarkUnmarshal700ROAR100Tags-32                  100          10079172 ns/op         4788000 B/op      75600 allocs/op
BenchmarkUnmarshal800ROAR100Tags-32                  100          11343767 ns/op         5472000 B/op      86400 allocs/op
BenchmarkUnmarshal900ROAR100Tags-32                  100          12928238 ns/op         6156000 B/op      97200 allocs/op
BenchmarkUnmarshal1000ROAR100Tags-32                 100          14004111 ns/op         6840000 B/op     108000 allocs/op
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
