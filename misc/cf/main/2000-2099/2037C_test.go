//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2037/C
// https://codeforces.com/problemset/status/2037/problem/C
func TestCF2037C(t *testing.T) {
	t.Log("Current test is [CF2037C]")
	testCases := [][2]string{
		{
			`2
			3
			8`,
			`-1
			1 8 7 3 6 2 4 5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2037C, testCases, 0)
}
