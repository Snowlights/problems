//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2050/C
// https://codeforces.com/problemset/status/2050/problem/C
func TestCF2050C(t *testing.T) {
	t.Log("Current test is [CF2050C]")
	testCases := [][2]string{
		{
			`9
			123
			322
			333333333333
			9997
			5472778912773
			1234567890
			23
			33
			52254522632`,
			`NO
			YES
			YES
			NO
			NO
			YES
			NO
			YES
			YES`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2050C, testCases, 0)
}
