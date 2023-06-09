# Go-wildcard

[![Go Report Card](https://goreportcard.com/badge/github.com/IGLOU-EU/go-wildcard)](https://goreportcard.com/report/github.com/IGLOU-EU/go-wildcard)
[![Go Reference](https://img.shields.io/badge/api-reference-blue)](https://pkg.go.dev/github.com/IGLOU-EU/go-wildcard)
[![BSD 3 Clause ](https://img.shields.io/badge/license-BSD_3_Clause-blue)](https://opensource.org/license/bsd-3-clause/)

## 💡 Why
The purpose of this library is to provide a simple and fast wildcard pattern matching.
Regex are much more complex, and slower (even prepared regex)... and the filepath.Match is not enough flexible.

So, this library is a very simple, very fast and a more flexible alternative to regex and filepath.Match. 
There are no dependencies and is alocation free. 🥳

## 🧰 Features
### Supported patterns operators
- `*` match zero or more characters
- `?` match zero or one character
- `.` match exactly one character

### Supported flags
- `FLAG_NONE` no flag
- `FLAG_CASEFOLD` ignore case

This is irrelevant for now, but you can combine them with `|` operator.   
For example: `FLAG_CASEFOLD|FLAG_NONE`

Because of the `strings.ToLower` operation, using `FLAG_CASEFOLD` flag is slower and result allocation...   
Even if this function is not self recursive, prefer to prepare your data before, because `strings.ToLower` is called for pattern and given string.

## 🧐 How to
>⚠️ WARNING: Unlike the GNU "libc", this library have no equivalent to "FNM_FILE_NAME". 
>To do this you can use "path/filepath" https://pkg.go.dev/path/filepath#Match

There is super simple to use this library, you just have to import it and use the Match function.
```go
package main

import (
	"fmt"

	wildcard "github.com/IGLOU-EU/go-wildcard"
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
- Match(t.pattern, t.name, FLAG_NONE) `The actual implementation`
- Match(t.pattern, t.name, FLAG_CASEFOLD) `The actual implementation, but with a strings.ToLower operation`

```bash
goos: linux
goarch: amd64
pkg: github.com/IGLOU-EU/go-wildcard
cpu: AMD Ryzen 7 PRO 6850U with Radeon Graphics     

BenchmarkRegex/0-16              1000000          1344 ns/op         766 B/op          9 allocs/op
BenchmarkRegex/1-16               115628         11343 ns/op        6592 B/op         26 allocs/op
BenchmarkRegex/2-16              5015937           263.6 ns/op       160 B/op          2 allocs/op
BenchmarkRegex/3-16               109844         13607 ns/op        6646 B/op         26 allocs/op
BenchmarkRegex/4-16               119311         13226 ns/op        7440 B/op         38 allocs/op
BenchmarkRegex/5-16              5427733           247.1 ns/op       160 B/op          2 allocs/op

BenchmarkFilepath/0-16          479149471            2.109 ns/op           0 B/op          0 allocs/op
BenchmarkFilepath/1-16           9473259           119.8 ns/op         0 B/op          0 allocs/op
BenchmarkFilepath/2-16          151451250            7.945 ns/op           0 B/op          0 allocs/op
BenchmarkFilepath/3-16           8295160           144.5 ns/op         0 B/op          0 allocs/op
BenchmarkFilepath/4-16          57564092            19.49 ns/op        0 B/op          0 allocs/op
BenchmarkFilepath/5-16           4911076           240.2 ns/op         0 B/op          0 allocs/op

BenchmarkOldMatchSimple/0-16    1000000000           0.4878 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatchSimple/1-16     4959806           283.4 ns/op       176 B/op          1 allocs/op
BenchmarkOldMatchSimple/2-16    1000000000           0.9326 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatchSimple/3-16     1910574           725.4 ns/op       352 B/op          2 allocs/op
BenchmarkOldMatchSimple/4-16     2918268           511.0 ns/op       336 B/op          2 allocs/op
BenchmarkOldMatchSimple/5-16    10512922           113.3 ns/op         0 B/op          0 allocs/op

BenchmarkOldMatch/0-16          1000000000           0.5195 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatch/1-16           5393407           279.9 ns/op       176 B/op          1 allocs/op
BenchmarkOldMatch/2-16          1000000000           0.7294 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatch/3-16           2123334           775.0 ns/op       352 B/op          2 allocs/op
BenchmarkOldMatch/4-16           2616628           496.1 ns/op       336 B/op          2 allocs/op
BenchmarkOldMatch/5-16          10818404           114.0 ns/op         0 B/op          0 allocs/op

BenchmarkMatch/0-16             454892713            2.281 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/1-16             349304181            2.928 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/2-16             476728009            2.127 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/3-16             290518609            3.822 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/4-16             124291632            8.709 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/5-16             23469135            48.56 ns/op        0 B/op          0 allocs/op

BenchmarkMatchCasefold/0-16      7673000           216.3 ns/op        48 B/op          1 allocs/op
BenchmarkMatchCasefold/1-16      7057550           223.6 ns/op        48 B/op          1 allocs/op
BenchmarkMatchCasefold/2-16      6580888           215.7 ns/op        48 B/op          1 allocs/op
BenchmarkMatchCasefold/3-16      3193930           451.2 ns/op        96 B/op          2 allocs/op
BenchmarkMatchCasefold/4-16      3213775           473.1 ns/op        96 B/op          2 allocs/op
BenchmarkMatchCasefold/5-16      2305406           481.4 ns/op        32 B/op          1 allocs/op

PASS
ok      github.com/IGLOU-EU/go-wildcard 60.533s
```

## 🕰 History 
Originally, this library was a fork from the Minio project.
The purpose was to give access to this "lib" under Apache license, without importing the entire Minio project.
And to keep it usable under the Apache License Version 2.0 after MinIO project is migrated to GNU Affero General Public License 3.0 or later from [`update license change for MinIO`](https://github.com/minio/minio/commit/069432566fcfac1f1053677cc925ddafd750730a)

The actual Minio wildcard matching code can be found in [`wildcard.go`](https://github.com/minio/pkg/tree/main/wildcard)