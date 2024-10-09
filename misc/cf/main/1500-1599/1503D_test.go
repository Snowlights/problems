//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1503/D
// https://codeforces.com/problemset/status/1503/problem/D
func TestCF1503D(t *testing.T) {
	t.Log("Current test is [CF1503D]")
	testCases := [][2]string{
		{
			`5
			3 10
			6 4
			1 9
			5 8
			2 7`,
			`2`,
		},
		{
			`2
			1 2
			3 4`,
			`-1`,
		},
		{
			`3
			1 2
			3 6
			4 5`,
			`-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1503D, testCases, 0)
}
