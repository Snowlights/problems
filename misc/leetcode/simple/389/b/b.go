package main

import "strings"

func countSubstrings(s string, c byte) int64 {
	k := int64(strings.Count(s, string(c)))
	return k * (k + 1) / 2
}
