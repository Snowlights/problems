//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2045/I
// https://codeforces.com/problemset/status/2045/problem/I
func TestCF2045I(t *testing.T) {
	t.Log("Current test is [CF2045I]")
	testCases := [][2]string{
		{
			`5 4
3 2 1 3 2`,
			`13`,
		},
		{
			`3 3
1 1 1`,
			`2`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF2045I, testCases, 0)
}
