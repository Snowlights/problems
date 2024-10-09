//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1689/C
// https://codeforces.com/problemset/status/1689/problem/C
func TestCF1689C(t *testing.T) {
	t.Log("Current test is [CF1689C]")
	testCases := [][2]string{
		{
			`4
			2
			1 2
			4
			1 2
			2 3
			2 4
			7
			1 2
			1 5
			2 3
			2 4
			5 6
			5 7
			15
			1 2
			2 3
			3 4
			4 5
			4 6
			3 7
			2 8
			1 9
			9 10
			9 11
			10 12
			10 13
			11 14
			11 15`,
			`0
			2
			2
			10
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1689C, testCases, 0)
}
