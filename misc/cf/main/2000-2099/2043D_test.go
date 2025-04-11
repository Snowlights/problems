//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2043/D
// https://codeforces.com/problemset/status/2043/problem/D
func TestCF2043D(t *testing.T) {
	t.Log("Current test is [CF2043D]")
	testCases := [][2]string{
		{
			`4
			4 8 2
			4 8 3
			4 8 4
			5 7 6`,
			`4 6
			-1 -1
			4 8
			6 6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2043D, testCases, 0)
}
