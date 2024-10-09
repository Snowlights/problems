//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1365/C
// https://codeforces.com/problemset/status/1365/problem/C
func TestCF1365C(t *testing.T) {
	t.Log("Current test is [CF1365C]")
	testCases := [][2]string{
		{
			`5
			1 2 3 4 5
			2 3 4 5 1`,
			`5`,
		},
		{
			`5
			5 4 3 2 1
			1 2 3 4 5
			`,
			`1`,
		},
		{
			`4
			1 3 2 4
			4 2 3 1
			`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1365C, testCases, 0)
}
