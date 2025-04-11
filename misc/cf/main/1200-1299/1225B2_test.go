//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1225/B2
// https://codeforces.com/problemset/status/1225/problem/B2
func TestCF1225B2(t *testing.T) {
	t.Log("Current test is [CF1225B2]")
	testCases := [][2]string{
		{
			`4
			5 2 2
			1 2 1 2 1
			9 3 3
			3 3 3 2 2 2 1 1 1
			4 10 4
			10 8 6 4
			16 9 8
			3 1 4 1 5 9 2 6 5 3 5 8 9 7 9 3`,
			`2
			1
			4
			5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1225B2, testCases, 0)
}
