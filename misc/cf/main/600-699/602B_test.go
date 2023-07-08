//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/602/B
// https://codeforces.com/problemset/status/602/problem/B
func TestCF602B(t *testing.T) {
	t.Log("Current test is [CF602B]")
	testCases := [][2]string{
		{
			`5
			1 2 3 3 2
			`,
			`4`,
		},
		{
			`11
			5 4 5 5 6 7 8 8 8 7 6
			`,
			`5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF602B, testCases, 0)
}
