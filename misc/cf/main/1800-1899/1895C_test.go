//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1895/C
// https://codeforces.com/problemset/status/1895/problem/C
func TestCF1895C(t *testing.T) {
	t.Log("Current test is [CF1895C]")
	testCases := [][2]string{
		{
			`10
			5 93746 59 3746 593 746 5937 46 59374 6`,
			`20`,
		},
		{
			`5
			2 22 222 2222 22222`,
			`13`,
		},
		{
			`3
			1 1 1`,
			`9`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1895C, testCases, 0)
}
