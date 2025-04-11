//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1632/B
// https://codeforces.com/problemset/status/1632/problem/B
func TestCF1632B(t *testing.T) {
	t.Log("Current test is [CF1632B]")
	testCases := [][2]string{
		{
			`4
			2
			3
			5
			10`,
			`0 1
			2 0 1
			3 2 1 0 4
			4 6 3 2 0 8 9 1 7 5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1632B, testCases, 0)
}
