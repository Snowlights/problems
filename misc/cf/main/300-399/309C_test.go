//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/309/C
// https://codeforces.com/problemset/status/309/problem/C
func TestCF309C(t *testing.T) {
	t.Log("Current test is [CF309C]")
	testCases := [][2]string{
		{
			`5 3
			8 4 3 2 2
			3 2 2
			`,
			`2`,
		},
		{
			`10 6
			1 1 1 1 1 1 1 1 1 1
			0 0 0 0 0 0
			`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF309C, testCases, 0)
}
