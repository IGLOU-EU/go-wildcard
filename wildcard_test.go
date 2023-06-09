/*
 * Copyright (c) 2023 Iglou.eu <contact@iglou.eu>
 * Copyright (c) 2023 Adrien Kara <adrien@iglou.eu>
 *
 * Licensed under the BSD 3-Clause License,
 * see LICENSE.md for more details.
 */

package wildcard_test

import (
	"testing"

	wildcard "github.com/IGLOU-EU/go-wildcard"
)

// TestMatch - Tests validate the logic of wild card matching.
// `Match` supports '*' and '?' wildcards.
// Sample usage: In resource matching for bucket policy validation.
func TestMatch(t *testing.T) {
	cases := []struct {
		s       string
		pattern string
		flag    wildcard.Flags
		result  bool
	}{
		{"", "", wildcard.FLAG_NONE, true},
		{"", "*", wildcard.FLAG_NONE, true},
		{"", "**", wildcard.FLAG_NONE, true},
		{"", "?", wildcard.FLAG_NONE, true},
		{"", "??", wildcard.FLAG_NONE, true},
		{"", "?*", wildcard.FLAG_NONE, true},
		{"", "*?", wildcard.FLAG_NONE, true},
		{"", ".", wildcard.FLAG_NONE, false},
		{"", ".?", wildcard.FLAG_NONE, false},
		{"", "?.", wildcard.FLAG_NONE, false},
		{"", ".*", wildcard.FLAG_NONE, false},
		{"", "*.", wildcard.FLAG_NONE, false},
		{"", "*.?", wildcard.FLAG_NONE, false},
		{"", "?.*", wildcard.FLAG_NONE, false},

		{"a", "", wildcard.FLAG_NONE, false},
		{"a", "a", wildcard.FLAG_NONE, true},
		{"a", "*", wildcard.FLAG_NONE, true},
		{"a", "**", wildcard.FLAG_NONE, true},
		{"a", "?", wildcard.FLAG_NONE, true},
		{"a", "??", wildcard.FLAG_NONE, true},
		{"a", ".", wildcard.FLAG_NONE, true},
		{"a", ".?", wildcard.FLAG_NONE, true},
		{"a", "?.", wildcard.FLAG_NONE, false},
		{"a", ".*", wildcard.FLAG_NONE, true},
		{"a", "*.", wildcard.FLAG_NONE, true},
		{"a", "*.?", wildcard.FLAG_NONE, true},
		{"a", "?.*", wildcard.FLAG_NONE, false},

		{"match the exact string", "match the exact string", wildcard.FLAG_NONE, true},
		{"do not match a different string", "this is a different string", wildcard.FLAG_NONE, false},
		{"Match The Exact String WITH DIFFERENT CASE", "Match The Exact String WITH DIFFERENT CASE", wildcard.FLAG_NONE, true},
		{"do not match a different string WITH DIFFERENT CASE", "this is a different string WITH DIFFERENT CASE", wildcard.FLAG_NONE, false},
		{"Do Not Match The Exact String With Different Case", "do not match the exact string with different case", wildcard.FLAG_NONE, false},
		{"match an emoji 😃", "match an emoji 😃", wildcard.FLAG_NONE, true},
		{"do not match because of different emoji 😃", "do not match because of different emoji 😄", wildcard.FLAG_NONE, false},
		{"🌅☕️📰👨‍💼👩‍💼🏢🖥️💼💻📊📈📉👨‍👩‍👧‍👦🍝🕰️💪🏋️‍♂️🏋️‍♀️🏋️‍♂️💼🚴‍♂️🚴‍♀️🚴‍♂️🛀💤🌃", "🌅☕️📰👨‍💼👩‍💼🏢🖥️💼💻📊📈📉👨‍👩‍👧‍👦🍝🕰️💪🏋️‍♂️🏋️‍♀️🏋️‍♂️💼🚴‍♂️🚴‍♀️🚴‍♂️🛀💤🌃", wildcard.FLAG_NONE, true},
		{"🌅☕️📰👨‍💼👩‍💼🏢🖥️💼💻📊📈📉👨‍👩‍👧‍👦🍝🕰️💪🏋️‍♂️🏋️‍♀️🏋️‍♂️💼🚴‍♂️🚴‍♀️🚴‍♂️🛀💤🌃", "🦌🐇🦡🐿️🌲🌳🏰🌳🌲🌞🌧️❄️🌬️⛈️🔥🎄🎅🎁🎉🎊🥳👨‍👩‍👧‍👦💏👪💖👩‍💼🛀", wildcard.FLAG_NONE, false},

		{"match a string with a *", "match a string *", wildcard.FLAG_NONE, true},
		{"match a string with a * at the beginning", "* at the beginning", wildcard.FLAG_NONE, true},
		{"match a string with two *", "match * with *", wildcard.FLAG_NONE, true},
		{"do not match a string with extra and a *", "do not match a string * with more", wildcard.FLAG_NONE, false},

		{"match a string with a ?", "match ? string with a ?", wildcard.FLAG_NONE, true},
		{"match a string with a ? at the beginning", "?atch a string with a ? at the beginning", wildcard.FLAG_NONE, true},
		{"match a string with two ?", "match a string with two ??", wildcard.FLAG_NONE, true},
		{"match a optional char with a ?", "match a optional? char with a ?", wildcard.FLAG_NONE, true},
		{"match a optional   char with a ?", "match a optional?   char with a ?", wildcard.FLAG_NONE, true},
		{"do not match a string with extra and a ?", "do not match ? string with extra and a ? like this", wildcard.FLAG_NONE, false},

		{"match a string with a .", "match . string with a .", wildcard.FLAG_NONE, true},
		{"match a string with a . at the beginning", ".atch a string with a . at the beginning", wildcard.FLAG_NONE, true},
		{"match a string with two .", "match a ..ring with two .", wildcard.FLAG_NONE, true},
		{"do not match a string with extra .", "do not match a string with extra ..", wildcard.FLAG_NONE, false},

		{"A big brown fox jumps over the lazy dog, with all there wildcards friends", ". big?brown fox jumps over * wildcard. friend??", wildcard.FLAG_NONE, true},
		{"A big brown fox fails to jump over the lazy dog, with all there wildcards friends", ". big?brown fox jumps over * wildcard. friend??", wildcard.FLAG_NONE, false},

		{"This IS a StrinG witH soMMe UppeRCase FriendS", "thIs is A stRINg wITh sOMmE uPpERcAse fRiENds", wildcard.FLAG_CASEFOLD, true},
	}

	for i, c := range cases {
		result := wildcard.Match(c.pattern, c.s, c.flag)
		if c.result != result {
			t.Errorf("Test %d: Expected `%v`, found `%v`; With Pattern: `%s` and String: `%s`", i+1, c.result, result, c.pattern, c.s)
		}
	}
}

func FuzzMatch(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if !wildcard.Match(s, s) {
			t.Fatalf("%s does not match %s", s, s)
		}
	})
}
