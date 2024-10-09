//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/761/D
// https://codeforces.com/problemset/status/761/problem/D
func TestCF761D(t *testing.T) {
	t.Log("Current test is [CF761D]")
	testCases := [][2]string{
		{
			`5 1 5
			1 1 1 1 1
			3 1 5 4 2`,
			`3 1 5 4 2`,
		},
		{
			`4 2 9
			3 4 8 9
			3 2 1 4`,
			`2 2 2 9`,
		},
		{
			`6 1 5
			1 1 1 1 1 1
			2 3 5 4 1 6`,
			`-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF761D, testCases, 0)
}
