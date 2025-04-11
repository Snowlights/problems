//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2028/C
// https://codeforces.com/problemset/status/2028/problem/C
func TestCF2028C(t *testing.T) {
	t.Log("Current test is [CF2028C]")
	testCases := [][2]string{
		{
			`7
			6 2 1
			1 1 10 1 1 10
			6 2 2
			1 1 10 1 1 10
			6 2 3
			1 1 10 1 1 10
			6 2 10
			1 1 10 1 1 10
			6 2 11
			1 1 10 1 1 10
			6 2 12
			1 1 10 1 1 10
			6 2 12
			1 1 1 1 10 10`,
			`22
			12
			2
			2
			2
			0
			-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2028C, testCases, 0)
}
