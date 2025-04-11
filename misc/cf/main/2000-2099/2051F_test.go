//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2051/F
// https://codeforces.com/problemset/status/2051/problem/F
func TestCF2051F(t *testing.T) {
	t.Log("Current test is [CF2051F]")
	testCases := [][2]string{
		{
			`5
			6 5 3
			1 2 3
			2 1 4
			2 1 1 2
			5 3 1
			3
			3 2 4
			2 1 1 1
			18 15 4
			13 15 1 16`,
			`2 3 5 
			2 2 2 2 
			2 
			2 3 3 3 
			2 4 6 8`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2051F, testCases, 0)
}
