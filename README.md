# Go-wilcard

[![Go Report Card](https://goreportcard.com/badge/git.iglou.eu/Imported/go-wilcard)](https://goreportcard.com/report/git.iglou.eu/Imported/go-wilcard)
[![Go Reference](https://img.shields.io/badge/api-reference-blue)](https://pkg.go.dev/git.iglou.eu/Imported/go-wilcard)
[![Go coverage](https://img.shields.io/badge/coverage-100%25-success)](https://img.shields.io)
[![Apache V2 License](https://img.shields.io/badge/license-Apache%202-blue)](https://opensource.org/licenses/MIT)

>Go-Wilcard is forked from [Minio project](https://github.com/minio/minio)   
>https://github.com/minio/minio/tree/master/pkg/wildcard

## Quick Example

This example shows a Go file which pattern matching ...  
You can use the Github repos to `github.com/IGLOU-EU/go-wilcard`
```go
package main

import (
	"fmt"
	"log"

	wildcard "git.iglou.eu/Imported/go-wilcard"
)

func main() {
    str := "daaadabadmanda"
    
    pattern := "da*da*da*"
    result := wildcard.Match(pattern, str)
	fmt.Println(str, pattern, result)

    pattern = "?a*da*d?*"
    result = wildcard.Match(pattern, str)
	fmt.Println(str, pattern, result)
}
```
