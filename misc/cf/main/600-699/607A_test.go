//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/607/A
// https://codeforces.com/problemset/status/607/problem/A
func TestCF607A(t *testing.T) {
	t.Log("Current test is [CF607A]")
	testCases := [][2]string{
		{
			`4
			1 9
			3 1
			6 1
			7 4
			`,
			`1`,
		},
		{
			`7
			1 1
			2 1
			3 1
			4 1
			5 1
			6 1
			7 1
			`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF607A, testCases, 0)
}
