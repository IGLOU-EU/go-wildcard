package wildcard_test

import (
	"fmt"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/IGLOU-EU/go-wildcard/v2"
)

var TestSet = []struct {
	pattern string
	name    string
}{
	{"", "These aren't the wildcard you're looking for"},
	{"These aren't the wildcard you're looking for", ""},
	{"*", "These aren't the wildcard you're looking for"},
	{"These aren't the wildcard you're looking for", "These aren't the wildcard you're looking for"},
	{"Th.e * the wildcard you?re looking fo?", "These aren't the wildcard you're looking for"},
	{"*🤷🏾‍♂️*", "T🥵🤷🏾‍♂️🥓"},
}

func BenchmarkRegex(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				regexp.MatchString(t.pattern, t.name)
			}
		})
	}
}

func BenchmarkFilepath(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				filepath.Match(t.pattern, t.name)
			}
		})
	}
}

func BenchmarkOldMatchSimple(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Old_MatchSimple(t.pattern, t.name)
			}
		})
	}
}

func BenchmarkOldMatch(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Old_Match(t.pattern, t.name)
			}
		})
	}
}

func BenchmarkMatch(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				wildcard.Match(t.pattern, t.name)
			}
		})
	}
}
