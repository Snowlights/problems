//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1139/C
// https://codeforces.com/problemset/status/1139/problem/C
func TestCF1139C(t *testing.T) {
	t.Log("Current test is [CF1139C]")
	testCases := [][2]string{
		{
			`4 4
			1 2 1
			2 3 1
			3 4 1
			`,
			`252`,
		},
		{
			`4 6
			1 2 0
			1 3 0
			1 4 0
			`,
			`0`,
		},
		{
			`3 5
			1 2 1
			2 3 0
			`,
			`210`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1139C, testCases, 0)
}
