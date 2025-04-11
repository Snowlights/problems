//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2033/D
// https://codeforces.com/problemset/status/2033/problem/D
func TestCF2033D(t *testing.T) {
	t.Log("Current test is [CF2033D]")
	testCases := [][2]string{
		{
			`3
			5
			2 1 -3 2 1
			7
			12 -4 4 43 -3 -5 8
			6
			0 -4 0 3 0 1`,
			`1
			2
			3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2033D, testCases, 0)
}
