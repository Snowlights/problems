//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1902/problem/C
// https://codeforces.com/problemset/status/1902/problem/C
func TestCF1902C(t *testing.T) {
	t.Log("Current test is [CF1902C]")
	testCases := [][2]string{
		{
			`3
			3
			1 2 3
			5
			1 -19 17 -3 -15
			1
			10`,
			`6
			27
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1902C, testCases, 0)
}
