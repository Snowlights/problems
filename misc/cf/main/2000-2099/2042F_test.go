//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2042/F
// https://codeforces.com/problemset/status/2042/problem/F
func TestCF2042F(t *testing.T) {
	t.Log("Current test is [CF2042F]")
	testCases := [][2]string{
		{
			`7
			3 -1 4 -3 2 4 0
			0 6 1 0 -3 -2 -1
			6
			3 1 7
			1 2 0
			3 3 6
			2 5 -3
			1 3 2
			3 1 5`,
			`18
			7
			16`,
		},
		{
			`10
			2 -1 -3 -2 0 4 5 6 2 5
			2 -4 -5 -1 6 2 5 -6 4 2
			10
			3 6 7
			1 10 -2
			3 5 7
			3 2 8
			2 1 -5
			2 7 4
			3 1 3
			3 3 8
			3 2 3
			1 4 4`,
			`23
			28
			28
			-17
			27
			-22`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2042F, testCases, 0)
}
