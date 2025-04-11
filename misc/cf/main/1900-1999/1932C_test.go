//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1932/C
// https://codeforces.com/problemset/status/1932/problem/C
func TestCF1932C(t *testing.T) {
	t.Log("Current test is [CF1932C]")
	testCases := [][2]string{
		{
			`4
			4 6
			3 1 4 2
			LRRL
			5 1
			1 1 1 1 1
			LLLLL
			6 8
			1 2 3 4 5 6
			RLLLRR
			1 10000
			10000
			R`,
			`0 2 4 1 
			0 0 0 0 0 
			0 0 0 4 4 4 
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1932C, testCases, 0)
}
