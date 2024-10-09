//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/721/problem/C
// https://codeforces.com/problemset/status/721/problem/C
func TestCF721C(t *testing.T) {
	t.Log("Current test is [CF721C]")
	testCases := [][2]string{
		{
			`4 3 13
			1 2 5
			2 3 7
			2 4 8
			`,
			`3
			1 2 4 
			`,
		},
		{
			`6 6 7
			1 2 2
			1 3 3
			3 6 3
			2 4 2
			4 6 2
			6 5 1
			`,
			`4
			1 2 4 6 
			`,
		},
		{
			`5 5 6
			1 3 3
			3 5 3
			1 2 2
			2 4 3
			4 5 2
			`,
			`3
			1 3 5 
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF721C, testCases, 0)
}
