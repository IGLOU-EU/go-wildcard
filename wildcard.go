/*
 * Copyright (c) 2023 Iglou.eu <contact@iglou.eu>
 * Copyright (c) 2023 Adrien Kara <adrien@iglou.eu>
 *
 * Licensed under the BSD 3-Clause License,
 * see LICENSE.md for more details.
 */

//go:generate go run cmd/build/build.go

// Package wildcard reports whether a string matches a wildcard pattern.
//
// Three operators are supported:
//   - '*' matches zero or more characters
//   - '?' matches zero or one character
//   - '.' matches exactly one character
//
// Any other character must match itself.
package wildcard

import "bytes"

// Match reports whether s matches pattern, comparing byte by byte and
// without allocating. Against multi-byte UTF-8 the operators apply to
// bytes, not whole characters; use MatchByRune when that matters.
func Match(pattern, s string) bool {
	if pattern == "" {
		return s == pattern
	}
	if pattern == "*" || s == pattern {
		return true
	}

	return matchByString(pattern, s)
}

// MatchByRune reports whether s matches pattern, comparing rune by rune,
// so the operators apply to whole Unicode code points. Converting pattern
// and s to runes allocates; prefer Match when byte semantics are enough.
func MatchByRune(pattern, s string) bool {
	if pattern == "" {
		return s == pattern
	}
	if pattern == "*" || s == pattern {
		return true
	}

	return matchByRunes([]rune(pattern), []rune(s))
}

// MatchFromByte is Match for byte slices: it reports whether s matches
// pattern, with the same byte-wise semantics and without allocation.
func MatchFromByte(pattern, s []byte) bool {
	if len(pattern) == 0 {
		return len(s) == 0
	}
	if string(pattern) == "*" || bytes.Equal(pattern, s) {
		return true
	}

	return matchByByte(pattern, s)
}
