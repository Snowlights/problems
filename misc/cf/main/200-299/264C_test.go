//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/264/C
// https://codeforces.com/problemset/status/264/problem/C
func TestCF264C(t *testing.T) {
	t.Log("Current test is [CF264C]")
	testCases := [][2]string{
		{
			`6 3
			1 -2 3 4 0 -1
			1 2 1 2 1 1
			5 1
			-2 1
			1 0`,
			`20`,
		},
		{
			`4 1
			-3 6 -1 2
			1 2 3 1
			1 -1`,
			`5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF264C, testCases, 0)
}
