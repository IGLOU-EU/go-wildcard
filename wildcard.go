/*
 * Copyright (c) 2023 Iglou.eu <contact@iglou.eu>
 * Copyright (c) 2023 Adrien Kara <adrien@iglou.eu>
 *
 * Licensed under the BSD 3-Clause License,
 * see LICENSE.md for more details.
 */

//go:generate go run cmd/build/build.go

package wildcard

import "bytes"

// Match returns true if the pattern matches the s string.
// The pattern can contain the wildcard characters '?' '.' and '*'.
func Match(pattern, s string) bool {
	if pattern == "" {
		return s == pattern
	}
	if pattern == "*" || s == pattern {
		return true
	}

	return matchByString(pattern, s)
}

func MatchByRune(pattern, s string) bool {
	if pattern == "" {
		return s == pattern
	}
	if pattern == "*" || s == pattern {
		return true
	}

	return matchByRunes([]rune(pattern), []rune(s))
}

func MatchFromByte(pattern, s []byte) bool {
	if len(pattern) == 0 {
		return len(s) == 0
	}
	if pattern[0] == '*' || bytes.Equal(pattern, s) {
		return true
	}

	return matchByByte(pattern, s)
}
