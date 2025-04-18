//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1466/D
// https://codeforces.com/problemset/status/1466/problem/D
func TestCF1466D(t *testing.T) {
	t.Log("Current test is [CF1466D]")
	testCases := [][2]string{
		{
			`4
			4
			3 5 4 6
			2 1
			3 1
			4 3
			2
			21 32
			2 1
			6
			20 13 17 13 13 11
			2 1
			3 1
			4 1
			5 1
			6 1
			4
			10 6 6 6
			1 2
			2 3
			4 1`,
			`18 22 25
			53
			87 107 127 147 167
			28 38 44`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1466D, testCases, 0)
}
