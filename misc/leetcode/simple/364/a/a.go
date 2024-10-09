package main

import "strings"

func maximumOddBinaryNumber(s string) string {
	cnt1 := strings.Count(s, "1")
	return strings.Repeat("1", cnt1-1) + strings.Repeat("0", len(s)-cnt1) + "1"
}
