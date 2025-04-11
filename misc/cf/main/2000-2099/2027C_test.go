//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2027/C
// https://codeforces.com/problemset/status/2027/problem/C
func TestCF2027C(t *testing.T) {
	t.Log("Current test is [CF2027C]")
	testCases := [][2]string{
		{
			`4
			5
			2 4 6 2 5
			5
			5 4 4 5 1
			4
			6 8 2 3
			1
			1`,
			`10
			11
			10
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2027C, testCases, 0)
}
