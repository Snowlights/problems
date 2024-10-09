//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1870/D
// https://codeforces.com/problemset/status/1870/problem/D
func TestCF1870D(t *testing.T) {
	t.Log("Current test is [CF1870D]")
	testCases := [][2]string{
		{
			`4
			3
			1 2 3
			5
			2
			3 4
			7
			3
			3 2 1
			2
			6
			10 6 4 6 3 4
			7`,
			`5 0 0 
			2 1 
			2 2 2 
			2 2 2 2 2 1 
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1870D, testCases, 0)
}
