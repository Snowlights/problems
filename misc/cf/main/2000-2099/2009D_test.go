//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2009/D
// https://codeforces.com/problemset/status/2009/problem/D
func TestCF2009D(t *testing.T) {
	t.Log("Current test is [CF2009D]")
	testCases := [][2]string{
		{
			`3
			5
			1 0
			1 1
			3 0
			5 0
			2 1
			3
			0 0
			1 0
			3 0
			9
			1 0
			2 0
			3 0
			4 0
			5 0
			2 1
			7 1
			8 1
			9 1`,
			`4
			0
			8`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2009D, testCases, 0)
}
