//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/817/B
// https://codeforces.com/problemset/status/817/problem/B
func TestCF817B(t *testing.T) {
	t.Log("Current test is [CF817B]")
	testCases := [][2]string{
		{
			`4
			1 1 1 1
			`,
			`4`,
		},
		{
			`5
			1 3 2 3 4
			`,
			`2`,
		},
		{
			`6
			1 3 3 1 3 2
			`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF817B, testCases, 0)
}
