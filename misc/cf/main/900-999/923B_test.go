//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/923/B
// https://codeforces.com/problemset/status/923/problem/B
func TestCF923B(t *testing.T) {
	t.Log("Current test is [CF923B]")
	testCases := [][2]string{
		{
			`3
			10 10 5
			5 7 2
			`,
			`5 12 4`,
		},
		{
			`5
			30 25 20 15 10
			9 10 12 4 13
			`,
			`9 20 35 11 25`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF923B, testCases, 0)
}
