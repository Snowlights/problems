//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/659/E
// https://codeforces.com/problemset/status/659/problem/E
func TestCF659E(t *testing.T) {
	t.Log("Current test is [CF659E]")
	testCases := [][2]string{
		{
			`4 3
			2 1
			1 3
			4 3
			`,
			`1`,
		},
		{
			`5 5
			2 1
			1 3
			2 3
			2 5
			4 3
			`,
			`0`,
		},
		{
			`6 5
			1 2
			2 3
			4 5
			4 6
			5 6
			`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF659E, testCases, 0)
}
