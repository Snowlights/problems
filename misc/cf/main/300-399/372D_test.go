//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/372/D
// https://codeforces.com/problemset/status/372/problem/D
func TestCF372D(t *testing.T) {
	t.Log("Current test is [CF372D]")
	testCases := [][2]string{
		{
			`10 6
			4 10
			10 6
			2 9
			9 6
			8 5
			7 1
			4 7
			7 3
			1 8`,
			`3`,
		},
		{
			`16 7
			13 11
			12 11
			2 14
			8 6
			9 15
			16 11
			5 14
			6 15
			4 3
			11 15
			15 14
			10 1
			3 14
			14 7
			1 7`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF372D, testCases, 0)
}
