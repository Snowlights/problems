//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/786/problem/C
// https://codeforces.com/problemset/status/786/problem/C
func TestCF786C(t *testing.T) {
	t.Log("Current test is [CF786C]")
	testCases := [][2]string{
		{
			`5
			1 3 4 3 3
			`,
			`4 2 1 1 1`,
		},
		{
			`8
			1 5 7 8 1 7 6 1
			`,
			`8 4 3 2 1 1 1 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF786C, testCases, 0)
}
