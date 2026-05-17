package wildcard

import "testing"

func TestMatchFromByteDoesNotMatchStarPrefixAlone(t *testing.T) {
	if MatchFromByte([]byte("*foo"), []byte("bar")) {
		t.Fatal("pattern *foo must not match bar")
	}
	if !MatchFromByte([]byte("*"), []byte("anything")) {
		t.Fatal("pattern * must match any string")
	}
}
