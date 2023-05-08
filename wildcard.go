/*
 * Copyright (c) 2023 Iglou.eu <contact@iglou.eu>
 * Copyright (c) 2023 Adrien Kara <adrien@iglou.eu>
 *
 * Licensed under the BSD 3-Clause License,
 * see LICENSE.md for more details.
 */

package wildcard

// Match returns true if the pattern matches the s string.
// The pattern can contain the wildcard characters '?' '.' and '*'.
func Match(pattern, s string) bool {
	if pattern == "" {
		return s == pattern
	}

	if pattern == "*" || s == pattern {
		return true
	}

	var patternIndex, sIndex, lastStar int
	patternLen := len(pattern)
	sLen := len(s)
	star := -1

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
	case '?', '.':
	case '*':
		star = patternIndex
		lastStar = sIndex
		patternIndex++
		goto Loop
	default:
		if pattern[patternIndex] != s[sIndex] {
			if star != -1 {
				patternIndex = star + 1
				lastStar++
				sIndex = lastStar
				goto Loop
			}
			return false
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
