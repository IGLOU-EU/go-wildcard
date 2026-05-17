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
// It supports the operators `*` (zero or more bytes), `?` (zero or one byte), and `.` (exactly one byte)
// using byte-wise comparison. For UTF-8 text where `?` or `.` must match a full rune, use MatchByRune.
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
// Same operators and byte-wise semantics as Match, without string conversion.
func MatchFromByte(pattern, s []byte) bool {
	if len(pattern) == 0 {
		return len(s) == 0
	}
	if pattern[0] == '*' || bytes.Equal(pattern, s) {
		return true
	}

	return matchByByte(pattern, s)
}
