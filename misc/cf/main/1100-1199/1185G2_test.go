//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1185/G2
// https://codeforces.com/problemset/status/1185/problem/G2
func TestCF1185G2(t *testing.T) {
	t.Log("Current test is [CF1185G2]")
	testCases := [][2]string{
		{
			`3 3
			1 1
			1 2
			1 3
			`,
			`6`,
		},
		{
			`3 3
			1 1
			1 1
			1 3
			`,
			`2`,
		},
		{
			`4 10
			5 3
			2 1
			3 2
			5 1
			`,
			`10`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1185G2, testCases, 0)
}
