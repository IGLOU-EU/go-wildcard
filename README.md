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
- Match(t.pattern, t.name) `The actual implementation`

```bash
goos: linux
goarch: amd64
pkg: github.com/IGLOU-EU/go-wildcard
cpu: AMD Ryzen 7 PRO 6850U with Radeon Graphics     
BenchmarkRegex/0-16              1000000          1322 ns/op         765 B/op          9 allocs/op
BenchmarkRegex/1-16               134851         10461 ns/op        6592 B/op         26 allocs/op
BenchmarkRegex/2-16              5871756           280.8 ns/op       160 B/op          2 allocs/op
BenchmarkRegex/3-16               108092         12096 ns/op        6647 B/op         26 allocs/op
BenchmarkRegex/4-16                92070         13924 ns/op        7436 B/op         38 allocs/op
BenchmarkRegex/5-16              4702372           277.6 ns/op       160 B/op          2 allocs/op

BenchmarkFilepath/0-16          548771120            1.836 ns/op           0 B/op          0 allocs/op
BenchmarkFilepath/1-16           9451810           117.8 ns/op         0 B/op          0 allocs/op
BenchmarkFilepath/2-16          151409767            7.853 ns/op           0 B/op          0 allocs/op
BenchmarkFilepath/3-16           8656650           143.8 ns/op         0 B/op          0 allocs/op
BenchmarkFilepath/4-16          67589983            18.33 ns/op        0 B/op          0 allocs/op
BenchmarkFilepath/5-16           4805623           240.3 ns/op         0 B/op          0 allocs/op

BenchmarkOldMatchSimple/0-16    1000000000           0.4971 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatchSimple/1-16     4738023           292.5 ns/op       176 B/op          1 allocs/op
BenchmarkOldMatchSimple/2-16    1000000000           0.9130 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatchSimple/3-16     1688683           763.8 ns/op       352 B/op          2 allocs/op
BenchmarkOldMatchSimple/4-16     2242758           514.0 ns/op       336 B/op          2 allocs/op
BenchmarkOldMatchSimple/5-16    10435084           110.7 ns/op         0 B/op          0 allocs/op

BenchmarkOldMatch/0-16          1000000000           0.4568 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatch/1-16           5300286           286.8 ns/op       176 B/op          1 allocs/op
BenchmarkOldMatch/2-16          1000000000           0.7127 ns/op          0 B/op          0 allocs/op
BenchmarkOldMatch/3-16           1608777           772.9 ns/op       352 B/op          2 allocs/op
BenchmarkOldMatch/4-16           2283015           548.9 ns/op       336 B/op          2 allocs/op
BenchmarkOldMatch/5-16          10425933           113.0 ns/op         0 B/op          0 allocs/op

BenchmarkMatch/0-16             654065395            1.774 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/1-16             352847413            2.973 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/2-16             652602918            1.822 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/3-16             412494770            2.940 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/4-16             197380323            5.447 ns/op           0 B/op          0 allocs/op
BenchmarkMatch/5-16             39741439            27.96 ns/op        0 B/op          0 allocs/op
PASS
ok      github.com/IGLOU-EU/go-wildcard 48.707s
```

## 🕰 History 
Originally, this library was a fork from the Minio project.
The purpose was to give access to this "lib" under Apache license, without importing the entire Minio project.
And to keep it usable under the Apache License Version 2.0 after MinIO project is migrated to GNU Affero General Public License 3.0 or later from [`update license change for MinIO`](https://github.com/minio/minio/commit/069432566fcfac1f1053677cc925ddafd750730a)

The actual Minio wildcard matching code can be found in [`wildcard.go`](https://github.com/minio/pkg/tree/main/wildcard)