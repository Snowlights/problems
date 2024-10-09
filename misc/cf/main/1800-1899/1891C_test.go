//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1891/C
// https://codeforces.com/problemset/status/1891/problem/C
func TestCF1891C(t *testing.T) {
	t.Log("Current test is [CF1891C]")
	testCases := [][2]string{
		{
			`4
			4
			1 3 1 1
			4
			1 2 1 1
			6
			3 2 1 5 2 4
			2
			1 6`,
			`4
			4
			11
			5
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1891C, testCases, 0)
}
