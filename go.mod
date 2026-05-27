// Deprecated: go-wildcard has moved to gitlab.com/iglou.eu/goulc/wildcard
// as part of the goulc library bundle. This module is no longer maintained;
// please migrate.
module github.com/IGLOU-EU/go-wildcard/v2

go 1.16

// MatchFromByte returned true for every pattern starting with '*'
// (e.g. "*abc" matched "xyz"). Fixed in v2.1.1 — please upgrade.
retract v2.1.0
