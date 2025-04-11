//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1324/F
// https://codeforces.com/problemset/status/1324/problem/F
func TestCF1324F(t *testing.T) {
	t.Log("Current test is [CF1324F]")
	testCases := [][2]string{
		{
			`9
			0 1 1 1 0 0 0 0 1
			1 2
			1 3
			3 4
			3 5
			2 6
			4 7
			6 8
			5 9`,
			`2 2 2 2 2 1 1 0 2`,
		},
		{
			`4
			0 0 1 0
			1 2
			1 3
			1 4`,
			`0 -1 1 -1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1324F, testCases, 0)
}
