//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2029/C
// https://codeforces.com/problemset/status/2029/problem/C
func TestCF2029C(t *testing.T) {
	t.Log("Current test is [CF2029C]")
	testCases := [][2]string{
		{
			`5
			6
			1 2 3 4 5 6
			7
			1 2 1 1 1 3 4
			1
			1
			9
			9 9 8 2 4 4 3 5 3
			10
			1 2 3 4 1 3 2 1 1 10`,
			`5
			4
			0
			4
			5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2029C, testCases, 0)
}
