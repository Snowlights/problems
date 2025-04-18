//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1736/C2
// https://codeforces.com/problemset/status/1736/problem/C2
func TestCF1736C2(t *testing.T) {
	t.Log("Current test is [CF1736C2]")
	testCases := [][2]string{
		{
			`4
			2 4 1 4
			3
			2 4
			3 3
			2 1`,
			`6
			10
			5`,
		},
		{
			`5
			1 1 3 2 1
			3
			1 3
			2 5
			4 5`,
			`7
			9
			8`,
		},
		{
			`10
			6 8 3 5 10 7 5 3 3 9
			1
			7 2`,
			`33`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1736C2, testCases, 0)
}
