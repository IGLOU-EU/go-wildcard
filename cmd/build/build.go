/*
 * Copyright (c) 2025 Iglou.eu <contact@iglou.eu>
 * Copyright (c) 2025 Adrien Kara <adrien@iglou.eu>
 *
 * Licensed under the BSD 3-Clause License,
 * see LICENSE.md for more details.
 */

package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"time"

	_ "embed"
)

type BuildArgs struct {
	FUNC_NAME           string
	COMPARISON_DOT      string
	COMPARISON_QUESTION string
	COMPARISON_STAR     string
	ARG_TYPE            string
	CLUSTER_TYPE        string
}

const (
	sourceFile = "source/wildcard_match.go"
	outputFile = "wildcard_match.go"
)

var (
	doNotEdit = []string{
		"// Code generated with go generate; DO NOT EDIT.",
		"// This file was generated by cmd/build/build.go at",
		"// " + time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST"),
		"// using source from " + sourceFile + "\n",
	}

	pkgName = "package " + os.Getenv("GOPACKAGE") + "\n\n"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	log.Println("Building wildcard_match.go")

	// Read the source
	source, err := os.ReadFile(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	// Type of function to build
	buildArgs := []BuildArgs{
		{
			FUNC_NAME:           "matchByString",
			COMPARISON_DOT:      "'.'",
			COMPARISON_QUESTION: "'?'",
			COMPARISON_STAR:     "'*'",
			ARG_TYPE:            "string",
			CLUSTER_TYPE:        "byte",
		},
		{
			FUNC_NAME:           "matchByByte",
			COMPARISON_DOT:      "'.'",
			COMPARISON_QUESTION: "'?'",
			COMPARISON_STAR:     "'*'",
			ARG_TYPE:            "[]byte",
			CLUSTER_TYPE:        "byte",
		},
		{
			FUNC_NAME:           "matchByRunes",
			COMPARISON_DOT:      "'.'",
			COMPARISON_QUESTION: "'?'",
			COMPARISON_STAR:     "'*'",
			ARG_TYPE:            "[]rune",
			CLUSTER_TYPE:        "rune",
		},
	}

	// Catch import and match function
	isBuild := ""
	var importBuilder strings.Builder
	var matchBuilder strings.Builder
	for _, line := range strings.Split(string(source), "\n") {
		if isBuild == "" {
			if strings.HasSuffix(line, "import (") {
				isBuild = "import"
			}
			if strings.HasPrefix(line, "func __FUNC_NAME__") {
				isBuild = "match"
			}
			if isBuild == "" {
				continue
			}
		}

		if isBuild == "import" {
			if line == ")" {
				isBuild = ""
			}

			importBuilder.WriteString(line + "\n")
			continue
		}

		if isBuild == "match" {
			matchBuilder.WriteString(line + "\n")
			continue
		}
	}

	log.Printf("Import size: %d\nMatch size: %d\n", importBuilder.Len(), matchBuilder.Len())

	// Build the output
	var output bytes.Buffer
	output.WriteString(strings.Join(doNotEdit, "\n"))
	output.WriteString(pkgName)
	output.WriteString(importBuilder.String())

	for _, args := range buildArgs {
		function := matchBuilder.String()
		function = strings.ReplaceAll(function, "__FUNC_NAME__", args.FUNC_NAME)
		function = strings.ReplaceAll(function, "__COMPARISON_DOT__", args.COMPARISON_DOT)
		function = strings.ReplaceAll(function, "__COMPARISON_QUESTION__", args.COMPARISON_QUESTION)
		function = strings.ReplaceAll(function, "__COMPARISON_STAR__", args.COMPARISON_STAR)
		function = strings.ReplaceAll(function, "__ARG_TYPE__", args.ARG_TYPE)
		function = strings.ReplaceAll(function, "__CLUSTER_TYPE__", args.CLUSTER_TYPE)

		output.WriteString(function)
	}

	// Save the output
	err = os.WriteFile(outputFile, output.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Output saved in " + outputFile + "\n")
}
