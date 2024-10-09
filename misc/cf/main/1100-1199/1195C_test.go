//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1195/C
// https://codeforces.com/problemset/status/1195/problem/C
func TestCF1195C(t *testing.T) {
	t.Log("Current test is [CF1195C]")
	testCases := [][2]string{
		{
			`5
			9 3 5 7 3
			5 8 1 4 5`,
			`29`,
		},
		{
			`3
			1 2 9
			10 1 1`,
			`19`,
		},
		{
			`1
			7
			4`,
			`7`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1195C, testCases, 0)
}
