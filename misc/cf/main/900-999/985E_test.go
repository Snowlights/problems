//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/985/E
// https://codeforces.com/problemset/status/985/problem/E
func TestCF985E(t *testing.T) {
	t.Log("Current test is [CF985E]")
	testCases := [][2]string{
		{
			`6 3 10
			7 2 7 7 4 2`,
			`YES`,
		},
		{
			`6 2 3
			4 5 3 13 4 10
			`,
			`YES`,
		},
		{
			`3 2 5
			10 16 22
			`,
			`NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF985E, testCases, 0)
}
