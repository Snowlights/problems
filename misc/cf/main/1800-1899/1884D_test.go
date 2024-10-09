//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1884/D
// https://codeforces.com/problemset/status/1884/problem/D
func TestCF1884D(t *testing.T) {
	t.Log("Current test is [CF1884D]")
	testCases := [][2]string{
		{
			`6
			4
			2 4 4 4
			4
			2 3 4 4
			9
			6 8 9 4 6 8 9 4 9
			9
			7 7 4 4 9 9 6 2 9
			18
			10 18 18 15 14 4 5 6 8 9 10 12 15 16 18 17 13 11
			21
			12 19 19 18 18 12 2 18 19 12 12 3 12 12 12 18 19 16 18 19 12`,
			`0
			3
			26
			26
			124
			82
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1884D, testCases, 0)
}
