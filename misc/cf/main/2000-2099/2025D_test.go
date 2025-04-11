//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2025/D
// https://codeforces.com/problemset/status/2025/problem/D
func TestCF2025D(t *testing.T) {
	t.Log("Current test is [CF2025D]")
	testCases := [][2]string{
		{
			`10 5
			0 1 0 2 0 -3 0 -4 0 -5`,
			`3`,
		},
		{
			`3 1
			1 -1 0`,
			`0`,
		},
		{
			`9 3
			0 0 1 0 2 -3 -2 -2 1`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2025D, testCases, 0)
}
