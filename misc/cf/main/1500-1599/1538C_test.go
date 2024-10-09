//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1538/C
// https://codeforces.com/problemset/status/1538/problem/C
func TestCF1538C(t *testing.T) {
	t.Log("Current test is [CF1538C]")
	testCases := [][2]string{
		{
			`4
			3 4 7
			5 1 2
			5 5 8
			5 1 2 4 3
			4 100 1000
			1 1 1 1
			5 9 13
			2 5 5 1 1
			`,
			`2
			7
			0
			1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1538C, testCases, 0)
}
