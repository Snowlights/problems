//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1343/C
// https://codeforces.com/problemset/status/1343/problem/C
func TestCF1343C(t *testing.T) {
	t.Log("Current test is [CF1343C]")
	testCases := [][2]string{
		{
			`4
			5
			1 2 3 -1 -2
			4
			-1 -2 -1 -3
			10
			-2 8 3 8 -4 -15 5 -2 -3 1
			6
			1 -1000000000 1 -1000000000 1 -1000000000`,
			`2
			-1
			6
			-2999999997
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1343C, testCases, 0)
}
