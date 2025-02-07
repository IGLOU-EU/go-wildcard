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

// Match returns true if the pattern matches the string s.
// It uses byte comparison rather than rune or grapheme cluster comparison.
// For matching complex Unicode, only the "*" wildcard or exact equality is supported.
func Match(pattern, s string) bool {
	if pattern == "" {
		return s == pattern
	}
	if pattern == "*" || s == pattern {
		return true
	}

	return matchByString(pattern, s)
}

// MatchByRune returns true if the pattern matches the string s.
// It supports complex Unicode matching with wildcards such as "*", "?", and ".".
// Note that it incurs allocation and more CPU usage.
func MatchByRune(pattern, s string) bool {
	if pattern == "" {
		return s == pattern
	}
	if pattern == "*" || s == pattern {
		return true
	}

	return matchByRunes([]rune(pattern), []rune(s))
}

// MatchFromByte returns true if the pattern matches the byte slice s.
// Similar to Match but operates on byte slices to avoid conversions/alloc.
func MatchFromByte(pattern, s []byte) bool {
	if len(pattern) == 0 {
		return len(s) == 0
	}
	if pattern[0] == '*' || bytes.Equal(pattern, s) {
		return true
	}

	return matchByByte(pattern, s)
}
