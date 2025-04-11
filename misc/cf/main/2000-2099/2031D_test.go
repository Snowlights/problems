//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2031/D
// https://codeforces.com/problemset/status/2031/problem/D
func TestCF2031D(t *testing.T) {
	t.Log("Current test is [CF2031D]")
	testCases := [][2]string{
		{
			`5
			4
			2 3 1 4
			5
			5 4 3 2 1
			4
			2 1 1 3
			4
			1 1 3 1
			8
			2 4 1 6 3 8 5 7`,
			`3 3 3 4 
			5 5 5 5 5 
			2 2 2 3 
			1 1 3 3 
			8 8 8 8 8 8 8 8`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2031D, testCases, 0)
}
