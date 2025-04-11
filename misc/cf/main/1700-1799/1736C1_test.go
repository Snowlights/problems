//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1736/C1
// https://codeforces.com/problemset/status/1736/problem/C1
func TestCF1736C1(t *testing.T) {
	t.Log("Current test is [CF1736C1]")
	testCases := [][2]string{
		{
			`3
			3
			1 2 3
			3
			1 1 1
			4
			2 1 4 3`,
			`6
			3
			7`,
		},
		{
			`1
			5
			1 3 1 1 1`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1736C1, testCases, 0)
}
