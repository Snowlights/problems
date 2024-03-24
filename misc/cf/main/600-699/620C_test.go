//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/620/C
// https://codeforces.com/problemset/status/620/problem/C
func TestCF620C(t *testing.T) {
	t.Log("Current test is [CF620C]")
	testCases := [][2]string{
		{
			`5
			1 2 3 4 1
			`,
			`1
			1 5`,
		},
		{
			`5
			1 2 3 4 5
			`,
			`-1`,
		},
		{
			`7
			1 2 1 3 1 2 1
			`,
			`2
			1 3
			4 7
			`,
		},
		{
			`10
			933677171 80672280 80672280 933677171 933677171 933677171 933677171 80672280 80672280 933677171`,
			`4
			1 3
			4 5
			6 7
			8 10`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF620C, testCases, 0)
}
