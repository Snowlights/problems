//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/91/B
// https://codeforces.com/problemset/status/91/problem/B
func TestCF91B(t *testing.T) {
	t.Log("Current test is [CF91B]")
	testCases := [][2]string{
		{
			`6
			10 8 5 3 50 45
			`,
			`2 1 0 -1 0 -1  `,
		},
		{
			`7
			10 4 6 3 2 8 15
			`,
			`4 2 1 0 -1 -1 -1`,
		},
		{
			`5
			10 3 1 10 11
			`,
			`1 0 -1 -1 -1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF91B, testCases, 0)
}
