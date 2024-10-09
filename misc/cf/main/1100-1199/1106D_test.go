//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1106/D
// https://codeforces.com/problemset/status/1106/problem/D
func TestCF1106D(t *testing.T) {
	t.Log("Current test is [CF1106D]")
	testCases := [][2]string{
		{
			`3 2
			1 2
			1 3
			`,
			`1 2 3`,
		},
		{
			`5 5
			1 4
			3 4
			5 4
			3 2
			1 5
			`,
			`1 4 3 2 5 
			`,
		},
		{
			`10 10
			1 4
			6 8
			2 5
			3 7
			9 4
			5 6
			3 4
			8 10
			8 9
			1 10
			`,
			`1 4 3 7 9 8 6 5 2 10 
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1106D, testCases, 0)
}
