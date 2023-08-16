# Go-wildcard

[![Go Report Card](https://goreportcard.com/badge/github.com/IGLOU-EU/go-wildcard/v2)](https://goreportcard.com/report/github.com/IGLOU-EU/go-wildcard/v2)
[![Go Reference](https://img.shields.io/badge/api-reference-blue)](https://pkg.go.dev/github.com/IGLOU-EU/go-wildcard/v2)
[![BSD 3 Clause ](https://img.shields.io/badge/license-BSD_3_Clause-blue)](https://opensource.org/license/bsd-3-clause/)

## 💡 Why
The purpose of this library is to provide a simple and fast wildcard pattern matching.
Regex are much more complex and slower (even prepared regex)... and the filepath.Match is file-name-centric.

So, this library is a very fast and very simple alternative to regex and not tied to filename semantics unlike filepath.Match. 
There are no dependencies and is alocation free. 🥳

## 🧰 Features
There are the supported patterns operators:
- `*` match zero or more characters
- `?` match zero or one character
- `.` match exactly one character

## 🧐 How to
>⚠️ WARNING: Unlike the GNU "libc", this library have no equivalent to "FNM_FILE_NAME". 
>To do this you can use "path/filepath" https://pkg.go.dev/path/filepath#Match

There is super simple to use this library, you just have to import it and use the Match function.
```go
package main

import (
	"fmt"

	"github.com/IGLOU-EU/go-wildcard/v2"
)

func main() {
    str := "daaadabadmanda"
    pattern := "?a*da*d.?*"

    result = wildcard.Match(pattern, str)
	fmt.Println(str, pattern, result)
}
```

## 🛸 Benchmark
The benchmark is done with the following command:
```bash
go test -benchmem -run=^$ -bench .
```

The tested fonctions are:
- regexp.MatchString(t.pattern, t.name)
- filepath.Match(t.pattern, t.name)
- oldMatchSimple(t.pattern, t.name) `From the commit a899be92514ed08aa5271bc3b93320b719ce2114`
- oldMatch(t.pattern, t.name) `From the commit a899be92514ed08aa5271bc3b93320b719ce2114`
- Match(t.pattern, t.name) `The actual implementation`

```bash
goos: linux
goarch: amd64
pkg: github.com/IGLOU-EU/go-wildcard/v2
cpu: AMD Ryzen 7 PRO 6850U with Radeon Graphics  

BenchmarkRegex/0-16              2062886           613.6 ns/op       767 B/op          9 allocs/op
BenchmarkRegex/1-16               240769          4891 ns/op        6592 B/op         26 allocs/op
BenchmarkRegex/2-16             11182353           106.1 ns/op       160 B/op          2 allocs/op
BenchmarkRegex/3-16               206820          5119 ns/op        6657 B/op         26 allocs/op
BenchmarkRegex/4-16               209696          5202 ns/op        7464 B/op         38 allocs/op
BenchmarkRegex/5-16             11510461           106.2 ns/op       160 B/op          2 allocs/op

BenchmarkFilepath/0-16          544894772            2.211 ns/op           0 B/op          0 allocs/op
BenchmarkFilepath/1-16           7447402           152.4 ns/op         0 B/op          0 allocs/op
BenchmarkFilepath/2-16          150307264            8.100 ns/op           0 B/op          0 allocs/op
BenchmarkFilepath/3-16           7204717           160.5 ns/op         0 B/op          0 allocs/op
BenchmarkFilepath/4-16          51796936            22.25 ns/op        0 B/op          0 allocs/op
BenchmarkFilepath/5-16           4137405           281.4 ns/op         0 B/op          0 allocs/op

BenchmarkOldMatchSimple/0-16    1000000000           0.5077 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatchSimple/1-16    10006898           140.4 ns/op       176 B/op          1 allocs/op
BenchmarkOldMatchSimple/2-16    1000000000           0.7710 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatchSimple/3-16     3102465           335.7 ns/op       352 B/op          2 allocs/op
BenchmarkOldMatchSimple/4-16     4941943           256.8 ns/op       336 B/op          2 allocs/op
BenchmarkOldMatchSimple/5-16     9047443           127.3 ns/op         0 B/op          0 allocs/op

BenchmarkOldMatch/0-16          1000000000           0.5054 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatch/1-16           9655593           142.0 ns/op       176 B/op          1 allocs/op
BenchmarkOldMatch/2-16          1000000000           0.9732 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatch/3-16           3008078           334.0 ns/op       352 B/op          2 allocs/op
BenchmarkOldMatch/4-16           4771064           251.4 ns/op       336 B/op          2 allocs/op
BenchmarkOldMatch/5-16           9545247           122.7 ns/op         0 B/op          0 allocs/op

BenchmarkMatch/0-16             1000000000           0.5314 ns/op          0 B/op          0 allocs/op
BenchmarkMatch/1-16             323842944            3.578 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/2-16             924416408            1.201 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/3-16             477003219            2.432 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/4-16             125328649            9.016 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/5-16             22776087            52.13 ns/op        0 B/op          0 allocs/op

PASS
ok      github.com/IGLOU-EU/go-wildcard/v2  39.328s
```

## 🕰 History 
Originally, this library was a fork from the Minio project.
The purpose was to give access to this "lib" under Apache license, without importing the entire Minio project.
And to keep it usable under the Apache License Version 2.0 after MinIO project is migrated to GNU Affero General Public License 3.0 or later from [`update license change for MinIO`](https://github.com/minio/minio/commit/069432566fcfac1f1053677cc925ddafd750730a)

The actual Minio wildcard matching code can be found in [`wildcard.go`](https://github.com/minio/pkg/tree/main/wildcard)