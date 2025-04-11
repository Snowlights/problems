//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2051/D
// https://codeforces.com/problemset/status/2051/problem/D
func TestCF2051D(t *testing.T) {
	t.Log("Current test is [CF2051D]")
	testCases := [][2]string{
		{
			`7
			4 8 10
			4 6 3 6
			6 22 27
			4 9 6 3 4 5
			3 8 10
			3 2 1
			3 1 1
			2 3 4
			3 3 6
			3 2 1
			4 4 12
			3 3 2 1
			6 8 8
			1 1 2 2 2 3`,
			`4
			7
			0
			0
			1
			5
			6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2051D, testCases, 0)
}
