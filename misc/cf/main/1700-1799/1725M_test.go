//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1725/M
// https://codeforces.com/problemset/status/1725/problem/M
func TestCF1725M(t *testing.T) {
	t.Log("Current test is [CF1725M]")
	testCases := [][2]string{
		{
			`5 7
			1 2 2
			2 4 1
			4 1 4
			2 5 3
			5 4 1
			5 2 4
			2 1 1`,
			`1 -1 3 4
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1725M, testCases, 0)
}
