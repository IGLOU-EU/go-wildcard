// Code generated with go generate; DO NOT EDIT.
// This file was generated by cmd/build/build.go at
// 2025-02-07 17:37:01.573285599 +0100 CET
// using source from source/wildcard_match.go
package wildcard

func matchByString(pattern, s string) bool {
	var lastErotemeCluster byte
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
		// '?' matches one character. Store its position and match exactly one character in the string.
		eroteme = patternIndex
		lastEroteme = sIndex
		lastErotemeCluster = byte(s[sIndex])
	case '*':
		// '*' matches zero or more characters. Store its position and increment the pattern index.
		star = patternIndex
		lastStar = sIndex
		patternIndex++
		goto Loop
	default:
		// If the characters don't match, check if there was a previous '?' or '*' to backtrack.
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

		// If the characters match, check if it was not the same to validate the eroteme.
		if eroteme != -1 && lastErotemeCluster != byte(s[sIndex]) {
			eroteme = -1
		}
	}

	patternIndex++
	sIndex++
	goto Loop

	// Check if the remaining pattern characters are '*' or '?', which can match the end of the string.
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

func matchByByte(pattern, s []byte) bool {
	var lastErotemeCluster byte
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
		// '?' matches one character. Store its position and match exactly one character in the string.
		eroteme = patternIndex
		lastEroteme = sIndex
		lastErotemeCluster = byte(s[sIndex])
	case '*':
		// '*' matches zero or more characters. Store its position and increment the pattern index.
		star = patternIndex
		lastStar = sIndex
		patternIndex++
		goto Loop
	default:
		// If the characters don't match, check if there was a previous '?' or '*' to backtrack.
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

		// If the characters match, check if it was not the same to validate the eroteme.
		if eroteme != -1 && lastErotemeCluster != byte(s[sIndex]) {
			eroteme = -1
		}
	}

	patternIndex++
	sIndex++
	goto Loop

	// Check if the remaining pattern characters are '*' or '?', which can match the end of the string.
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

func matchByRunes(pattern, s []rune) bool {
	var lastErotemeCluster rune
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
		// '?' matches one character. Store its position and match exactly one character in the string.
		eroteme = patternIndex
		lastEroteme = sIndex
		lastErotemeCluster = rune(s[sIndex])
	case '*':
		// '*' matches zero or more characters. Store its position and increment the pattern index.
		star = patternIndex
		lastStar = sIndex
		patternIndex++
		goto Loop
	default:
		// If the characters don't match, check if there was a previous '?' or '*' to backtrack.
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

		// If the characters match, check if it was not the same to validate the eroteme.
		if eroteme != -1 && lastErotemeCluster != rune(s[sIndex]) {
			eroteme = -1
		}
	}

	patternIndex++
	sIndex++
	goto Loop

	// Check if the remaining pattern characters are '*' or '?', which can match the end of the string.
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

