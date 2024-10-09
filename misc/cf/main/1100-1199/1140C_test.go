//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1140/C
// https://codeforces.com/problemset/status/1140/problem/C
func TestCF1140C(t *testing.T) {
	t.Log("Current test is [CF1140C]")
	testCases := [][2]string{
		{
			`4 3
			4 7
			15 1
			3 6
			6 8
			`,
			`78`,
		},
		{
			`5 3
			12 31
			112 4
			100 100
			13 55
			55 50
			`,
			`10000
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1140C, testCases, 0)
}
