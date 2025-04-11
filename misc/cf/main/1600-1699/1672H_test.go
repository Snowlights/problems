//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1672/H
// https://codeforces.com/problemset/status/1672/problem/H
func TestCF1672H(t *testing.T) {
	t.Log("Current test is [CF1672H]")
	testCases := [][2]string{
		{
			`5 3
			11011
			2 4
			1 5
			3 5`,
			`1
			3
			2`,
		},
		{
			`10 3
			1001110110
			1 10
			2 5
			5 10`,
			`4
			2
			3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1672H, testCases, 0)
}
