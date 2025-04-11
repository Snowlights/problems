//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2044/D
// https://codeforces.com/problemset/status/2044/problem/D
func TestCF2044D(t *testing.T) {
	t.Log("Current test is [CF2044D]")
	testCases := [][2]string{
		{
			`4
			2
			1 2
			4
			1 1 1 2
			8
			4 5 5 5 1 1 2 1
			10
			1 1 2 2 1 1 3 3 1 1`,
			`1 2
			1 1 2 2
			4 5 5 1 1 2 2 3
			1 8 2 2 1 3 3 9 1 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2044D, testCases, 0)
}
