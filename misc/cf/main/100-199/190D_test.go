//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/190/D
// https://codeforces.com/problemset/status/190/problem/D
func TestCF190D(t *testing.T) {
	t.Log("Current test is [CF190D]")
	testCases := [][2]string{
		{
			`4 2
			1 2 1 2
			`,
			`3`,
		},
		{
			`5 3
			1 2 1 1 3
			`,
			`2`,
		},
		{
			`3 1
			1 1 1
			`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF190D, testCases, 0)
}
