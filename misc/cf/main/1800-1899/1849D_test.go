//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1849/D
// https://codeforces.com/problemset/status/1849/problem/D
func TestCF1849D(t *testing.T) {
	t.Log("Current test is [CF1849D]")
	testCases := [][2]string{
		{
			`3
			0 2 0
			`,
			`1`,
		},
		{
			`4
			0 0 1 1
			`,
			`2`,
		},
		{
			`7
			0 1 0 0 1 0 2
			`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1849D, testCases, 0)
}
