//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/812/C
// https://codeforces.com/problemset/status/812/problem/C
func TestCF812C(t *testing.T) {
	t.Log("Current test is [CF812C]")
	testCases := [][2]string{
		{
			`3 11
			2 3 5`,
			`2 11`,
		},
		{
			`4 100
			1 2 5 6`,
			`4 54`,
		},
		{
			`1 7
			7`,
			`0 0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF812C, testCases, 0)
}
