//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1840/problem/F
// https://codeforces.com/problemset/status/1840/problem/F
func TestCF1840F(t *testing.T) {
	t.Log("Current test is [CF1840F]")
	testCases := [][2]string{
		{
			`5
			1 3
			4
			1 2 0
			2 2 1
			3 2 2
			4 1 1
			3 3
			6
			2 1 0
			2 1 1
			2 1 2
			2 2 0
			2 2 1
			2 2 2
			2 1
			3
			7 1 2
			2 1 1
			7 2 1
			2 2
			5
			9 1 2
			3 2 0
			5 1 2
			4 2 2
			7 1 0
			4 6
			7
			6 1 2
			12 1 3
			4 1 0
			17 2 3
			1 2 6
			16 2 6
			3 2 4`,
			`5
			-1
			3
			6
			10
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1840F, testCases, 0)
}
