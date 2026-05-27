# Go-wildcard

> [!WARNING]
> **Deprecated — moved to [`gitlab.com/iglou.eu/goulc/wildcard`](https://gitlab.com/iglou.eu/goulc).**
> This library has been merged into the `goulc` bundle that regroups several
> Iglou libraries under one module. This repository is no longer maintained;
> please migrate.

[![Go Report Card](https://goreportcard.com/badge/github.com/IGLOU-EU/go-wildcard/v2)](https://goreportcard.com/report/github.com/IGLOU-EU/go-wildcard/v2)
[![Go Reference](https://img.shields.io/badge/api-reference-blue)](https://pkg.go.dev/github.com/IGLOU-EU/go-wildcard/v2)
[![BSD 3 Clause ](https://img.shields.io/badge/license-BSD_3_Clause-blue)](https://opensource.org/license/bsd-3-clause/)

## 💡 Why
The purpose of this library is to provide a simple and fast wildcard pattern matching.
Regex are much more complex and slower (even prepared regex)... and the filepath.Match is file-name-centric.

So, this library is a very fast and very simple alternative to regex and not tied to filename semantics unlike filepath.Match.
There are no dependencies and is allocation-free. 🥳

## 🧰 Features
These are the supported pattern operators:
- `*` match zero or more characters
- `?` match zero or one character
- `.` match exactly one character

## 🧐 How to
>⚠️ WARNING: Unlike the GNU "libc", this library have no equivalent to "FNM_FILE_NAME". 
>To do this you can use "path/filepath" https://pkg.go.dev/path/filepath#Match

It is very simple to use this library: import it and call the Match function — or one of its variants, as shown below.
```go
package main

import (
	"fmt"

	"github.com/IGLOU-EU/go-wildcard/v2"
)

func main() {
	str := "daaadabadmanda"
	pattern := "?a*da*d.?*"

	resultM := wildcard.Match(pattern, str) // Fastest, compares byte by byte
	resultMFB := wildcard.MatchFromByte([]byte(pattern), []byte(str)) // Like Match, for byte slices. Skips the string conversion
	resultMBR := wildcard.MatchByRune(pattern, str) // Compares rune by rune. Slower and the []rune conversion allocates

	fmt.Println(str, pattern, resultM, resultMFB, resultMBR)
}
```

## 🛸 Benchmark
The benchmark is done with the following command:
```bash
go test -benchmem -bench . github.com/IGLOU-EU/go-wildcard/v2/benchmark
```

```yml
goos: linux
goarch: amd64
pkg: github.com/IGLOU-EU/go-wildcard/v2
cpu: AMD Ryzen 7 PRO 6850U with Radeon Graphics  
```

The tested functions are:
- regexp.MatchString
- regexp.MatchPreparedString
- filepath.Match
- oldMatchSimple `From the commit a899be92514ed08aa5271bc3b93320b719ce2114`
- oldMatch `From the commit a899be92514ed08aa5271bc3b93320b719ce2114`
- Match `From string with byte comparison`
- MatchByRune `From string with rune comparison`
- MatchFromByte `From byte slice with byte comparison`

<!-- BENCHMARK_TABLE:START -->

| Rank | Benchmark | Average ns/op | Samples |
| ---: | --- | ---: | ---: |
| 1 | BenchmarkMatch | 10.83 | 6 |
| 2 | BenchmarkMatchFromByte | 12.14 | 6 |
| 3 | BenchmarkMatchByRune | 60.78 | 6 |
| 4 | BenchmarkRegexPrepared | 86.48 | 4 |
| 5 | BenchmarkFilepath | 87.90 | 6 |
| 6 | BenchmarkOldMatch | 97.04 | 6 |
| 7 | BenchmarkOldMatchSimple | 98.24 | 6 |
| 8 | BenchmarkRegex | 2261.77 | 6 |

<!-- BENCHMARK_TABLE:END -->

![time bench](./assets/graph_time.png)
![allocs bench](./assets/graph_allocs.png)

## 🕰 History 
Originally, this library was a fork from the Minio project.
The purpose was to give access to this "lib" under Apache license, without importing the entire Minio project.
And to keep it usable under the Apache License Version 2.0 after MinIO project is migrated to GNU Affero General Public License 3.0 or later from [`update license change for MinIO`](https://github.com/minio/minio/commit/069432566fcfac1f1053677cc925ddafd750730a)

The actual Minio wildcard matching code can be found in [`wildcard.go`](https://github.com/minio/pkg/tree/main/wildcard)
