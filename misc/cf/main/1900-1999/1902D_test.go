//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1902/D
// https://codeforces.com/problemset/status/1902/problem/D
func TestCF1902D(t *testing.T) {
	t.Log("Current test is [CF1902D]")
	testCases := [][2]string{
		{
			`8 3
			RDLLUURU
			-1 2 1 7
			0 0 3 4
			0 1 7 8`,
			`YES
			YES
			NO`,
		},
		{
			`4 2
			RLDU
			0 0 2 2
			-1 -1 2 3`,
			`YES
			NO`,
		},
		{
			`10 6
			DLUDLRULLD
			-1 0 1 10
			-1 -2 2 5
			-4 -2 6 10
			-1 0 3 9
			0 1 4 7
			-3 -1 5 8`,
			`YES
			YES
			YES
			NO
			YES
			YES`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1902D, testCases, 0)
}
