//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2050/F
// https://codeforces.com/problemset/status/2050/problem/F
func TestCF2050F(t *testing.T) {
	t.Log("Current test is [CF2050F]")
	testCases := [][2]string{
		{
			`3
			5 5
			5 14 2 6 3
			4 5
			1 4
			2 4
			3 5
			1 1
			1 1
			7
			1 1
			3 2
			1 7 8
			2 3
			1 2`,
			`3 1 4 1 0 
			0 
			1 6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2050F, testCases, 0)
}
