package wildcard_bench

import (
	"fmt"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/IGLOU-EU/go-wildcard/v2"
)

var TestSet = []struct {
	pattern string
	input   string
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
				regexp.MatchString(t.pattern, t.input)
			}
		})
	}
}

func BenchmarkFilepath(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				filepath.Match(t.pattern, t.input)
			}
		})
	}
}

func BenchmarkOldMatchSimple(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Old_MatchSimple(t.pattern, t.input)
			}
		})
	}
}

func BenchmarkOldMatch(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Old_Match(t.pattern, t.input)
			}
		})
	}
}

func BenchmarkMatch(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				wildcard.Match(t.pattern, t.input)
			}
		})
	}
}

func BenchmarkMatchByRune(b *testing.B) {
	for i, t := range TestSet {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				wildcard.MatchByRune(t.pattern, t.input)
			}
		})
	}
}

func BenchmarkMatchFromByte(b *testing.B) {
	for i, t := range TestSet {
		pattern := []byte(t.pattern)
		input := []byte(t.input)

		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				wildcard.MatchFromByte(pattern, input)
			}
		})
	}
}
