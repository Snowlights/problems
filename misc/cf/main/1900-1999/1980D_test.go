//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1980/D
// https://codeforces.com/problemset/status/1980/problem/D
func TestCF1980D(t *testing.T) {
	t.Log("Current test is [CF1980D]")
	testCases := [][2]string{
		{
			`12
			6
			20 6 12 3 48 36
			4
			12 6 3 4
			3
			10 12 3
			5
			32 16 8 4 2
			5
			100 50 2 10 20
			4
			2 4 8 1
			10
			7 4 6 2 4 5 1 4 2 8
			7
			5 9 6 8 5 9 2
			6
			11 14 8 12 9 3
			9
			5 7 3 10 6 3 12 6 3
			3
			4 2 4
			8
			1 6 11 12 6 12 3 6`,
			`YES
			NO
			YES
			NO
			YES
			YES
			NO
			YES
			YES
			YES
			YES
			YES`,
		},
		{
			`1
			4
			6 2 3 5`,
			`YES`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1980D, testCases, 0)
}
