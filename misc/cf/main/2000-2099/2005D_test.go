//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2005/D
// https://codeforces.com/problemset/status/2005/problem/D
func TestCF2005D(t *testing.T) {
	t.Log("Current test is [CF2005D]")
	testCases := [][2]string{
		{
			`5
			8
			11 4 16 17 3 24 25 8
			8 10 4 21 17 18 25 21
			4
			6 4 24 13
			15 3 1 14
			2
			13 14
			5 8
			8
			20 17 15 11 21 10 3 7
			9 9 4 20 14 9 13 1
			2
			18 13
			15 20`,
			`2 36
			3 2
			2 3
			2 36
			6 1`,
		},
		{
			`1
			2
			5 5
			4 4`,
			`9 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2005D, testCases, 0)
}
