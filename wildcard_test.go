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
		result  bool
	}{
		{"", "", true},
		{"", "*", true},
		{"", "**", true},
		{"", "?", true},
		{"", "??", true},
		{"", "?*", true},
		{"", "*?", true},
		{"", ".", false},
		{"", ".?", false},
		{"", "?.", false},
		{"", ".*", false},
		{"", "*.", false},
		{"", "*.?", false},
		{"", "?.*", false},

		{"a", "", false},
		{"a", "a", true},
		{"a", "*", true},
		{"a", "**", true},
		{"a", "?", true},
		{"a", "??", true},
		{"a", ".", true},
		{"a", ".?", true},
		{"a", "?.", false},
		{"a", ".*", true},
		{"a", "*.", true},
		{"a", "*.?", true},
		{"a", "?.*", false},

		{"match the exact string", "match the exact string", true},
		{"do not match a different string", "this is a different string", false},
		{"Match The Exact String WITH DIFFERENT CASE", "Match The Exact String WITH DIFFERENT CASE", true},
		{"do not match a different string WITH DIFFERENT CASE", "this is a different string WITH DIFFERENT CASE", false},
		{"Do Not Match The Exact String With Different Case", "do not match the exact string with different case", false},
		{"match an emoji 😃", "match an emoji 😃", true},
		{"do not match because of different emoji 😃", "do not match because of different emoji 😄", false},
		{"🌅☕️📰👨‍💼👩‍💼🏢🖥️💼💻📊📈📉👨‍👩‍👧‍👦🍝🕰️💪🏋️‍♂️🏋️‍♀️🏋️‍♂️💼🚴‍♂️🚴‍♀️🚴‍♂️🛀💤🌃", "🌅☕️📰👨‍💼👩‍💼🏢🖥️💼💻📊📈📉👨‍👩‍👧‍👦🍝🕰️💪🏋️‍♂️🏋️‍♀️🏋️‍♂️💼🚴‍♂️🚴‍♀️🚴‍♂️🛀💤🌃", true},
		{"🌅☕️📰👨‍💼👩‍💼🏢🖥️💼💻📊📈📉👨‍👩‍👧‍👦🍝🕰️💪🏋️‍♂️🏋️‍♀️🏋️‍♂️💼🚴‍♂️🚴‍♀️🚴‍♂️🛀💤🌃", "🦌🐇🦡🐿️🌲🌳🏰🌳🌲🌞🌧️❄️🌬️⛈️🔥🎄🎅🎁🎉🎊🥳👨‍👩‍👧‍👦💏👪💖👩‍💼🛀", false},

		{"match a string with a *", "match a string *", true},
		{"match a string with a * at the beginning", "* at the beginning", true},
		{"match a string with two *", "match * with *", true},
		{"do not match a string with extra and a *", "do not match a string * with more", false},

		{"match a string with a ?", "match ? string with a ?", true},
		{"match a string with a ? at the beginning", "?atch a string with a ? at the beginning", true},
		{"match a string with two ?", "match a string with two ??", true},
		{"match a optional char with a ?", "match a optional? char with a ?", true},
		{"do not match a string with extra and a ?", "do not match ? string with extra and a ? like this", false},

		{"match a string with a .", "match . string with a .", true},
		{"match a string with a . at the beginning", ".atch a string with a . at the beginning", true},
		{"match a string with two .", "match a ..ring with two .", true},
		{"do not match a string with extra .", "do not match a string with extra ..", false},

		{"A big brown fox jumps over the lazy dog, with all there wildcards friends", ". big?brown fox jumps over * wildcard. friend??", true},
		{"A big brown fox fails to jump over the lazy dog, with all there wildcards friends", ". big?brown fox jumps over * wildcard. friend??", false},
	}

	for i, c := range cases {
		result := wildcard.Match(c.pattern, c.s)
		if c.result != result {
			t.Errorf("Test %d: Expected `%v`, found `%v`; With Pattern: `%s` and String: `%s`", i+1, c.result, result, c.pattern, c.s)
		}
	}
}
