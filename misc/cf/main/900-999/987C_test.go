//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/987/C
// https://codeforces.com/problemset/status/987/problem/C
func TestCF987C(t *testing.T) {
	t.Log("Current test is [CF987C]")
	testCases := [][2]string{
		{
			`5
			2 4 5 4 10
			40 30 20 10 40`,
			`90`,
		},
		{
			`3
			100 101 100
			2 4 5`,
			`-1`,
		},
		{
			`10
			1 2 3 4 5 6 7 8 9 10
			10 13 11 14 15 12 13 13 18 13`,
			`33`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF987C, testCases, 0)
}
