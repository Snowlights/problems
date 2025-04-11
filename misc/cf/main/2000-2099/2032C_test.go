//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2032/C
// https://codeforces.com/problemset/status/2032/problem/C
func TestCF2032C(t *testing.T) {
	t.Log("Current test is [CF2032C]")
	testCases := [][2]string{
		{
			`4
			7
			1 2 3 4 5 6 7
			3
			1 3 2
			3
			4 5 3
			15
			9 3 8 1 6 5 3 8 2 1 4 2 9 4 7`,
			`3
			1
			0
			8`,
		},
		{
			`1
			4
			1 1 1 2`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2032C, testCases, 0)
}
