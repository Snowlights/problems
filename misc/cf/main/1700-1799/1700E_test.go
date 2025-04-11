//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1700/E
// https://codeforces.com/problemset/status/1700/problem/E
func TestCF1700E(t *testing.T) {
	t.Log("Current test is [CF1700E]")
	testCases := [][2]string{
		{
			`3 3
			2 1 3
			6 7 4
			9 8 5`,
			`0`,
		},
		{
			`2 3
			1 6 4
			3 2 5`,
			`1 3`,
		},
		{
			`1 6
			1 6 5 4 3 2`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1700E, testCases, 0)
}
