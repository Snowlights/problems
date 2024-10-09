//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/487/B
// https://codeforces.com/problemset/status/487/problem/B
func TestCF487B(t *testing.T) {
	t.Log("Current test is [CF487B]")
	testCases := [][2]string{
		{
			`7 2 2
			1 3 1 2 4 1 2
			`,
			`3`,
		},
		{
			`7 2 2
			1 100 1 100 1 100 1
			`,
			`-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF487B, testCases, 0)
}
