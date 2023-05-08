/*
 * Copyright (c) 2023 Iglou.eu <contact@iglou.eu>
 * Copyright (c) 2023 Adrien Kara <adrien@iglou.eu>
 *
 * Licensed under the BSD 3-Clause License,
 * see LICENSE.md for more details.
 */

package wildcard

import "strings"

// Flags type is used to specify matching options for the Match function
type Flags uint8

// Constants definition for matching options
const (
	FLAG_NONE     = 1 << iota // No special behavior
	FLAG_CASEFOLD             // Case-insensitive match
)

// Match function checks if the given string s matches the wildcard pattern
// with specified matching options (Flags).
//
// Supported wildcards:
// `*` match zero or more characters
// `?` match zero or one character
// `.` match exactly one character
//
// Supported matching options:
// FLAG_NONE     - No special behavior
// FLAG_CASEFOLD - Case-insensitive match
func Match(pattern, s string, option Flags) bool {
	// If FLAG_CASEFOLD is set, convert both pattern and string to lowercase
	if option&FLAG_CASEFOLD != 0 {
		s = strings.ToLower(s)
		pattern = strings.ToLower(pattern)
	}

	if pattern == "" {
		return s == pattern
	}

	if pattern == "*" || s == pattern {
		return true
	}

	var lastErotemeByte byte
	var patternIndex, sIndex, lastStar, lastEroteme int
	patternLen := len(pattern)
	sLen := len(s)
	star := -1
	eroteme := -1

Loop:
	if sIndex >= sLen {
		goto checkPattern
	}

	if patternIndex >= patternLen {
		if star != -1 {
			patternIndex = star + 1
			lastStar++
			sIndex = lastStar
			goto Loop
		}
		return false
	}
	switch pattern[patternIndex] {
	case '.':
		// It matches any single character. So, we don't need to check anything.
	case '?':
		eroteme = patternIndex
		lastEroteme = sIndex
		lastErotemeByte = s[sIndex]
	case '*':
		star = patternIndex
		lastStar = sIndex
		patternIndex++
		goto Loop
	default:
		if pattern[patternIndex] != s[sIndex] {
			if eroteme != -1 {
				patternIndex = eroteme + 1
				sIndex = lastEroteme
				eroteme = -1
				goto Loop
			}

			if star != -1 {
				patternIndex = star + 1
				lastStar++
				sIndex = lastStar
				goto Loop
			}

			return false
		}

		if eroteme != -1 && lastErotemeByte != s[sIndex] {
			eroteme = -1
		}
	}

	patternIndex++
	sIndex++
	goto Loop

checkPattern:
	if patternIndex < patternLen {
		if pattern[patternIndex] == '*' {
			patternIndex++
			goto checkPattern
		} else if pattern[patternIndex] == '?' {
			if sIndex >= sLen {
				sIndex--
			}
			patternIndex++
			goto checkPattern
		}
	}

	return patternIndex == patternLen
}
