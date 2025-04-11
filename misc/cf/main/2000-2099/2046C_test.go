//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2046/C
// https://codeforces.com/problemset/status/2046/problem/C
func TestCF2046C(t *testing.T) {
	t.Log("Current test is [CF2046C]")
	testCases := [][2]string{
		{
			`4
			4
			1 1
			1 2
			2 1
			2 2
			4
			0 0
			0 0
			0 0
			0 0
			8
			1 2
			2 1
			2 -1
			1 -2
			-1 -2
			-2 -1
			-2 1
			-1 2
			7
			1 1
			1 2
			1 3
			1 4
			2 1
			3 1
			4 1`,
			`1
			2 2
			0
			0 0
			2
			1 0
			0
			0 0`,
		},
		{
			`1
			14
			-3 -7
			6 -4
			5 -8
			10 6
			-1 -10
			10 -1
			2 -5
			-9 -7
			-10 -4
			4 -7
			5 7
			-8 4
			6 -5
			-1 8`,
			`3
			2 -4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2046C, testCases, 0)
}
