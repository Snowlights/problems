//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1798/D
// https://codeforces.com/problemset/status/1798/problem/D
func TestCF1798D(t *testing.T) {
	t.Log("Current test is [CF1798D]")
	testCases := [][2]string{
		{
			`7
			4
			3 4 -2 -5
			5
			2 2 2 -3 -3
			8
			-3 -3 1 1 1 1 1 1
			3
			0 1 -1
			7
			-3 4 3 4 -4 -4 0
			1
			0
			7
			-18 13 -18 -17 12 15 13`,
			`Yes
			-5 -2 3 4
			Yes
			-3 2 -3 2 2
			Yes
			1 1 1 -3 1 1 1 -3
			Yes
			-1 0 1
			Yes
			4 -4 4 -4 0 3 -3
			No
			Yes
			13 12 -18 15 -18 13 -17`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1798D, testCases, 0)
}
