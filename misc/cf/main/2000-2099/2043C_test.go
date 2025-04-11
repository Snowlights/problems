//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2043/C
// https://codeforces.com/problemset/status/2043/problem/C
func TestCF2043C(t *testing.T) {
	t.Log("Current test is [CF2043C]")
	testCases := [][2]string{
		{
			`5
			5
			1 -1 10 1 1
			5
			-1 -1 -1 -1 -1
			2
			-1 2
			2
			7 1
			3
			1 4 -1`,
			`8
			-1 0 1 2 9 10 11 12 
			6
			-5 -4 -3 -2 -1 0 
			4
			-1 0 1 2 
			4
			0 1 7 8 
			6
			-1 0 1 3 4 5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2043C, testCases, 0)
}
