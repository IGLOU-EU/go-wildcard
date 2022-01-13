# Go-wildcard

[![Go Report Card](https://goreportcard.com/badge/git.iglou.eu/Imported/go-wildcard)](https://goreportcard.com/report/git.iglou.eu/Imported/go-wildcard)
[![Go Reference](https://img.shields.io/badge/api-reference-blue)](https://pkg.go.dev/git.iglou.eu/Imported/go-wildcard)
[![Go coverage](https://gocover.io/_badge/git.iglou.eu/Imported/go-wildcard)](https://gocover.io/git.iglou.eu/Imported/go-wildcard)
[![Apache V2 License](https://img.shields.io/badge/license-Apache%202-blue)](https://opensource.org/licenses/MIT)

>Go-wildcard is forked from [Minio project](https://github.com/minio/minio)   
>https://github.com/minio/minio/tree/master/pkg/wildcard

## Why
This part of Minio project is a very cool, fast and light wildcard pattern matching.   
But using it, need to import the full Minio project inside your own ... And this is a cool, but BIG project.   

Two function are available `MatchSimple` and `Match`   
- `MatchSimple` only covert `*` usage (he is a bit faster)
- `Match` support full wildcard matching, `*` and `?`

I know Regex, but this is a big part, and it is slow (even prepared regex) ...   
I know Glob, but most of the time, I only need simple wildcard matching.   

## How to
Using GitHub repo
```sh
go get github.com/IGLOU-EU/go-wildcard@latest
```

## Quick Example

This example shows a Go file which pattern matching ...  
```go
package main

import (
	"fmt"
	"log"

	wildcard "github.com/IGLOU-EU/go-wildcard"
)

func main() {
    str := "daaadabadmanda"
    
    pattern := "da*da*da*"
    result := wildcard.MatchSimple(pattern, str)
	fmt.Println(str, pattern, result)

    pattern = "?a*da*d?*"
    result = wildcard.Match(pattern, str)
	fmt.Println(str, pattern, result)
}
```
