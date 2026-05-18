module github.com/IGLOU-EU/go-wildcard/v2

go 1.16

// MatchFromByte returned true for every pattern starting with '*'
// (e.g. "*abc" matched "xyz"). Fixed in v2.1.1 — please upgrade.
retract v2.1.0
