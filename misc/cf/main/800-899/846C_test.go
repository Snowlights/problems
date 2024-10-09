//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/846/problem/C
// https://codeforces.com/problemset/status/846/problem/C
func TestCF846C(t *testing.T) {
	t.Log("Current test is [CF846C]")
	testCases := [][2]string{
		{
			`3
			-1 2 3
			`,
			`0 1 3`,
		},
		{
			`4
			0 0 -1 0
			`,
			`0 0 0`,
		},
		{
			`1
			10000
			`,
			`1 1 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF846C, testCases, 0)
}
