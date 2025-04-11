//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2027/D1
// https://codeforces.com/problemset/status/2027/problem/D1
func TestCF2027D1(t *testing.T) {
	t.Log("Current test is [CF2027D1]")
	testCases := [][2]string{
		{
			`5
			4 2
			9 3 4 3
			11 7
			1 2
			20
			19 18
			10 2
			2 5 2 1 10 3 2 9 9 6
			17 9
			10 11
			2 2 2 2 2 2 2 2 2 2
			20 18 16 14 12 10 8 6 4 2 1
			1 6
			10
			32 16 8 4 2 1`,
			`1
			-1
			2
			10
			4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2027D1, testCases, 0)
}
