//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/213/C
// https://codeforces.com/problemset/status/213/problem/C
func TestCF213C(t *testing.T) {
	t.Log("Current test is [CF213C]")
	testCases := [][2]string{
		{
			`1
			5
			`,
			`5`,
		},
		{
			`2
			11 14
			16 12
			`,
			`53`,
		},
		{
			`3
			25 16 25
			12 18 19
			11 13 8
			`,
			`136`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF213C, testCases, 0)
}
