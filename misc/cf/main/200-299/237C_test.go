//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/237/problem/C
// https://codeforces.com/problemset/status/237/problem/C
func TestCF237C(t *testing.T) {
	t.Log("Current test is [CF237C]")
	testCases := [][2]string{
		{
			`2 4 2`,
			`3`,
		},
		{
			`6 13 1`,
			`4`,
		},
		{
			`1 4 3`,
			`-1`,
		},
		{
			`8 10 3`,
			`-1`,
		},
		{
			`4 10 2`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF237C, testCases, -1)
}
