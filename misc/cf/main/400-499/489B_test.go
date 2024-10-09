//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/489/B
// https://codeforces.com/problemset/status/489/problem/B
func TestCF489B(t *testing.T) {
	t.Log("Current test is [CF489B]")
	testCases := [][2]string{
		{
			`4
			1 4 6 2
			5
			5 1 5 7 9
			`,
			`3`,
		},
		{
			`4
			1 2 3 4
			4
			10 11 12 13
			`,
			`0`,
		},
		{
			`5
			1 1 1 1 1
			3
			1 2 3
			`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF489B, testCases, 0)
}
