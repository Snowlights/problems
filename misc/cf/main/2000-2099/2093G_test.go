//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2093/G
// https://codeforces.com/problemset/status/2093/problem/G
func TestCF2093G(t *testing.T) {
	t.Log("Current test is [CF2093G]")
	testCases := [][2]string{
		{
			`6
			5 0
			1 2 3 4 5
			5 7
			1 2 3 4 5
			5 8
			1 2 3 4 5
			5 7
			3 5 1 4 2
			5 3
			3 5 1 4 2
			6 71
			26 56 12 45 60 27`,
			`1
			2
			-1
			4
			2
			-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2093G, testCases, 0)
}
