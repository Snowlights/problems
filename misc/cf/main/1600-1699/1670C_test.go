//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1670/C
// https://codeforces.com/problemset/status/1670/problem/C
func TestCF1670C(t *testing.T) {
	t.Log("Current test is [CF1670C]")
	testCases := [][2]string{
		{
			`9
			7
			1 2 3 4 5 6 7
			2 3 1 7 6 5 4
			2 0 1 0 0 0 0
			1
			1
			1
			0
			6
			1 5 2 4 6 3
			6 5 3 1 4 2
			6 0 0 0 0 0
			8
			1 6 4 7 2 3 8 5
			3 2 8 1 4 5 6 7
			1 0 0 7 0 3 0 5
			10
			1 8 6 2 4 7 9 3 10 5
			1 9 2 3 4 10 8 6 7 5
			1 9 2 3 4 10 8 6 7 5
			7
			1 2 3 4 5 6 7
			2 3 1 7 6 5 4
			0 0 0 0 0 0 0
			5
			1 2 3 4 5
			1 2 3 4 5
			0 0 0 0 0
			5
			1 2 3 4 5
			1 2 3 5 4
			0 0 0 0 0
			3
			1 2 3
			3 1 2
			0 0 0`,
			`4
			1
			2
			2
			1
			8
			1
			2
			2
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1670C, testCases, 0)
}
