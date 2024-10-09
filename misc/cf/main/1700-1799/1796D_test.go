//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1796/D
// https://codeforces.com/problemset/status/1796/problem/D
func TestCF1796D(t *testing.T) {
	t.Log("Current test is [CF1796D]")
	testCases := [][2]string{
		{
			`4
			4 1 2
			2 -1 2 3
			2 2 3
			-1 2
			3 0 5
			3 2 4
			6 2 -8
			4 -1 9 -3 7 -8`,
			`5
			7
			0
			44`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1796D, testCases, 0)
}
