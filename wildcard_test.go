/*
 * Copyright (c) 2026 Iglou.eu <contact@iglou.eu>
 * Copyright (c) 2026 Adrien Kara <adrien@iglou.eu>
 *
 * Licensed under the BSD 3-Clause License,
 * see LICENSE.md for more details.
 */

package wildcard

import "testing"

// TestMatch only covers the trivial cases handled inline by the public API
// functions. Each case is run against Match, MatchByRune and MatchFromByte,
// which must all agree. The wildcard matching logic itself is tested in the
// source/ package.
func TestMatch(t *testing.T) {
	cases := []struct {
		name    string
		pattern string
		s       string
		want    bool
	}{
		{"empty pattern and empty string", "", "", true},
		{"empty pattern and non-empty string", "", "something", false},
		{"lone star matches anything", "*", "anything", true},
		{"lone star matches empty string", "*", "", true},
		{"exact equality", "exact", "exact", true},
		{"pattern starting with a star but not matching", "*abc", "xyz", false},
		{"delegates to matcher on match", "ex?ct", "exact", true},
		{"delegates to matcher on mismatch", "ex?ct", "different", false},
	}

	matchers := []struct {
		name string
		fn   func(pattern, s string) bool
	}{
		{"Match", Match},
		{"MatchByRune", MatchByRune},
		{"MatchFromByte", func(pattern, s string) bool {
			return MatchFromByte([]byte(pattern), []byte(s))
		}},
	}

	for _, c := range cases {
		for _, m := range matchers {
			if got := m.fn(c.pattern, c.s); got != c.want {
				t.Errorf("%s: %s(%q, %q) = %v, want %v",
					c.name, m.name, c.pattern, c.s, got, c.want)
			}
		}
	}
}
