//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2033/C
// https://codeforces.com/problemset/status/2033/problem/C
func TestCF2033C(t *testing.T) {
	t.Log("Current test is [CF2033C]")
	testCases := [][2]string{
		{
			`9
			5
			1 1 1 2 3
			6
			2 1 2 2 1 1
			4
			1 2 1 1
			6
			2 1 1 2 2 4
			4
			2 1 2 3
			6
			1 2 2 1 2 1
			5
			4 5 5 1 5
			7
			1 4 3 5 1 1 3
			7
			3 1 3 2 2 3 3`,
			`1
			2
			1
			0
			0
			1
			1
			0
			2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2033C, testCases, 0)
}
