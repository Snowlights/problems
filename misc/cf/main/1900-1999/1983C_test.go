//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1983/C
// https://codeforces.com/problemset/status/1983/problem/C
func TestCF1983C(t *testing.T) {
	t.Log("Current test is [CF1983C]")
	testCases := [][2]string{
		{
			`10
			5
			5 1 1 1 1
			1 1 5 1 1
			1 1 1 1 5
			6
			1 2 3 4 5 6
			5 6 1 2 3 4
			3 4 5 6 1 2
			4
			4 4 4 4
			4 4 4 4
			4 4 4 4
			5
			5 10 5 2 10
			9 6 9 7 1
			10 7 10 2 3
			3
			4 5 2
			6 1 4
			1 8 2
			3
			10 4 10
			8 7 9
			10 4 10
			7
			57113 65383 19795 53580 74452 3879 23255
			12917 16782 89147 93107 27365 15044 43095
			33518 63581 33565 34112 46774 44151 41756
			6
			6 3 1 8 7 1
			10 2 6 2 2 4
			10 9 2 1 2 2
			5
			5 5 4 5 5
			1 6 3 8 6
			2 4 1 9 8
			10
			1 1 1 1 1001 1 1 1001 1 1
			1 1 1 1 1 1 2001 1 1 1
			1 1 1 1 1 1001 1 1 1 1001`,
			`1 1 2 3 4 5 
			5 6 1 2 3 4 
			-1
			-1
			1 1 3 3 2 2 
			-1
			1 2 3 4 5 7 
			3 6 1 1 2 2 
			1 2 3 4 5 5 
			1 5 6 7 8 10`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1983C, testCases, 0)
}
