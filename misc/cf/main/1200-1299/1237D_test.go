//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1237/D
// https://codeforces.com/problemset/status/1237/problem/D
func TestCF1237D(t *testing.T) {
	t.Log("Current test is [CF1237D]")
	testCases := [][2]string{
		{
			`4
			11 5 2 7`,
			`1 1 3 2`,
		},
		{
			`4
			3 2 5 3`,
			`5 4 3 6`,
		},
		{
			`3
			4 3 6`,
			`-1 -1 -1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1237D, testCases, 0)
}
