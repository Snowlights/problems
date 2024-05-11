//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1623/C
// https://codeforces.com/problemset/status/1623/problem/C
func TestCF1623C(t *testing.T) {
	t.Log("Current test is [CF1623C]")
	testCases := [][2]string{
		{
			`4
			4
			1 2 10 100
			4
			100 100 100 1
			5
			5 1 1 1 8
			6
			1 2 3 4 5 6`,
			`7
			1
			1
			3
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1623C, testCases, 0)
}
