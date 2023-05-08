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
