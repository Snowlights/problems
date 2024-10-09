//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1401/F
// https://codeforces.com/problemset/status/1401/problem/F
func TestCF1401F(t *testing.T) {
	t.Log("Current test is [CF1401F]")
	testCases := [][2]string{
		{
			`2 3
			7 4 9 9
			1 2 8
			3 1
			4 2 4`,
			`24`,
		},
		{
			`3 8
			7 0 8 8 7 1 5 2
			4 3 7
			2 1
			3 2
			4 1 6
			2 3
			1 5 16
			4 8 8
			3 0`,
			`29
			22
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1401F, testCases, 0)
}
