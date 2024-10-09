//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/864/D
// https://codeforces.com/problemset/status/864/problem/D
func TestCF864D(t *testing.T) {
	t.Log("Current test is [CF864D]")
	testCases := [][2]string{
		{
			`4
			3 2 2 3`,
			`2
			1 2 4 3`,
		},
		{
			`6
			4 5 6 3 2 1`,
			`0
			4 5 6 3 2 1 `,
		},
		{
			`10
			6 8 4 6 7 1 6 3 4 5`,
			`3
			2 8 4 6 7 1 9 3 10 5 `,
		},
	}
	codeforces.AssertEqualStringCase(t, CF864D, testCases, 0)
}
