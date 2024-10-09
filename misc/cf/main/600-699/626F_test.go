//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/626/problem/F
// https://codeforces.com/problemset/status/626/problem/F
func TestCF626F(t *testing.T) {
	t.Log("Current test is [CF626F]")
	testCases := [][2]string{
		{
			`3 2
			2 4 5`,
			`3`,
		},
		{
			`4 3
			7 8 9 10`,
			`13`,
		},
		{
			`4 0
			5 10 20 21`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF626F, testCases, 0)
}
