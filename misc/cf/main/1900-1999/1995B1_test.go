//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1995/B1
// https://codeforces.com/problemset/status/1995/problem/B1
func TestCF1995B1(t *testing.T) {
	t.Log("Current test is [CF1995B1]")
	testCases := [][2]string{
		{
			`5
			5 10
			1 1 2 2 3
			8 20
			4 2 7 5 6 1 1 1
			8 100000
			239 30 610 122 24 40 8 2
			11 13
			2 4 11 1 1 2 3 5 4 3 2
			8 1033
			206 206 206 207 207 207 207 1000`,
			`7
			13
			610
			13
			1033`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1995B1, testCases, 0)
}
