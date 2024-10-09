//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/480/problem/C
// https://codeforces.com/problemset/status/480/problem/C
func TestCF480C(t *testing.T) {
	t.Log("Current test is [CF480C]")
	testCases := [][2]string{
		{
			`5 2 4 1`,
			`2`,
		},
		{
			`5 2 4 2`,
			`2`,
		},
		{
			`5 3 4 1`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF480C, testCases, 0)
}
