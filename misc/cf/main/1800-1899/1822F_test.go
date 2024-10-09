//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1822/F
// https://codeforces.com/problemset/status/1822/problem/F
func TestCF1822F(t *testing.T) {
	t.Log("Current test is [CF1822F]")
	testCases := [][2]string{
		{
			`4
			3 2 3
			2 1
			3 1
			5 4 1
			2 1
			4 2
			5 4
			3 4
			6 5 3
			4 1
			6 1
			2 6
			5 1
			3 2
			10 6 4
			1 3
			1 9
			9 7
			7 6
			6 4
			9 2
			2 8
			8 5
			5 10`,
			`2
			12
			17
			32
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1822F, testCases, 0)
}
