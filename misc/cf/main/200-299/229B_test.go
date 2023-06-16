//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/229/problem/B
// https://codeforces.com/problemset/status/229/problem/B
func TestCF229B(t *testing.T) {
	t.Log("Current test is [CF229B]")
	testCases := [][2]string{
		{
			`4 6
			1 2 2
			1 3 3
			1 4 8
			2 3 4
			2 4 5
			3 4 3
			0
			1 3
			2 3 4
			0
			`,
			`7`,
		},
		{
			`3 1
			1 2 3
			0
			1 3
			0
			`,
			`-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF229B, testCases, 0)
}
