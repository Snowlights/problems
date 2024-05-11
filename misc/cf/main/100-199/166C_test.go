//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/166/C
// https://codeforces.com/problemset/status/166/problem/C
func TestCF166C(t *testing.T) {
	t.Log("Current test is [CF166C]")
	testCases := [][2]string{
		{
			`3 10
			10 20 30
			`,
			`1`,
		},
		{
			`3 4
			1 2 3
			`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF166C, testCases, 0)
}
