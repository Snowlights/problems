//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2040/C
// https://codeforces.com/problemset/status/2040/problem/C
func TestCF2040C(t *testing.T) {
	t.Log("Current test is [CF2040C]")
	testCases := [][2]string{
		{
			`6
			3 2
			3 3
			4 11
			4 6
			6 39
			7 34`,
			`1 3 2 
			2 3 1 
			-1
			2 4 3 1 
			-1
			2 3 4 5 7 6 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2040C, testCases, 0)
}
