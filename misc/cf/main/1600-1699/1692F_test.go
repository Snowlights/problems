//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1692/F
// https://codeforces.com/problemset/status/1692/problem/F
func TestCF1692F(t *testing.T) {
	t.Log("Current test is [CF1692F]")
	testCases := [][2]string{
		{
			`6
			4
			20 22 19 84
			4
			1 11 1 2022
			4
			1100 1100 1100 1111
			5
			12 34 56 78 90
			4
			1 9 8 4
			6
			16 38 94 25 18 99`,
			`YES
			YES
			NO
			NO
			YES
			YES
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1692F, testCases, 0)
}
