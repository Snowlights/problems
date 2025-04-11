//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2037/G
// https://codeforces.com/problemset/status/2037/problem/G
func TestCF2037G(t *testing.T) {
	t.Log("Current test is [CF2037G]")
	testCases := [][2]string{
		{
			`5
			2 6 3 4 6`,
			`5`,
		},
		{
			`5
			4 196 2662 2197 121`,
			`2`,
		},
		{
			`7
			3 6 8 9 11 12 20`,
			`7`,
		},
		{
			`2
			2 3`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2037G, testCases, 0)
}
