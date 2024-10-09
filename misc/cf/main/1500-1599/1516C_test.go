//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1516/problem/C
// https://codeforces.com/problemset/status/1516/problem/C
func TestCF1516C(t *testing.T) {
	t.Log("Current test is [CF1516C]")
	testCases := [][2]string{
		{
			`4
			6 3 9 12`,
			`1
			2`,
		},
		{
			`2
			1 2`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1516C, testCases, 0)
}
