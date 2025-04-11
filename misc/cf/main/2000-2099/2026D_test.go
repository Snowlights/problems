//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2026/D
// https://codeforces.com/problemset/status/2026/problem/D
func TestCF2026D(t *testing.T) {
	t.Log("Current test is [CF2026D]")
	testCases := [][2]string{
		{
			`4
			1 2 5 10
			15
			1 1
			1 2
			1 3
			1 4
			1 5
			1 10
			5 10
			6 10
			2 8
			3 4
			3 10
			3 8
			5 6
			5 5
			1 8`,
			`1
			4
			12
			30
			32
			86
			56
			54
			60
			26
			82
			57
			9
			2
			61`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2026D, testCases, 0)
}
