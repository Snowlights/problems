//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2069/C
// https://codeforces.com/problemset/status/2069/problem/C
func TestCF2069C(t *testing.T) {
	t.Log("Current test is [CF2069C]")
	testCases := [][2]string{
		{
			`4
			7
			3 2 1 2 2 1 3
			4
			3 1 2 2
			3
			1 2 3
			9
			1 2 3 2 1 3 2 2 3`,
			`3
			0
			1
			22`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2069C, testCases, 0)
}
