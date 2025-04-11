//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2042/C
// https://codeforces.com/problemset/status/2042/problem/C
func TestCF2042C(t *testing.T) {
	t.Log("Current test is [CF2042C]")
	testCases := [][2]string{
		{
			`7
			4 1
			1001
			4 1
			1010
			4 1
			0110
			4 2
			0110
			6 3
			001110
			10 20
			1111111111
			5 11
			11111`,
			`2
			-1
			2
			-1
			3
			4
			-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2042C, testCases, 0)
}
