//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/898/D
// https://codeforces.com/problemset/status/898/problem/D
func TestCF898D(t *testing.T) {
	t.Log("Current test is [CF898D]")
	testCases := [][2]string{
		{
			`3 3 2
			3 5 1
			`,
			`1`,
		},
		{
			`5 10 3
			12 8 18 25 1
			`,
			`0`,
		},
		{
			`7 7 2
			7 3 4 1 6 5 2
			`,
			`6`,
		},
		{
			`2 2 2
			1 3
			`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF898D, testCases, 0)
}
