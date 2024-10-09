//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1700/C
// https://codeforces.com/problemset/status/1700/problem/C
func TestCF1700C(t *testing.T) {
	t.Log("Current test is [CF1700C]")
	testCases := [][2]string{
		{
			`4
			3
			-2 -2 -2
			3
			10 4 7
			4
			4 -4 4 -4
			5
			1 -2 3 -4 5
			`,
			`2
			13
			36
			33
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1700C, testCases, 0)
}
