//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1793/D
// https://codeforces.com/problemset/status/1793/problem/D
func TestCF1793D(t *testing.T) {
	t.Log("Current test is [CF1793D]")
	testCases := [][2]string{
		{
			`3
			1 3 2
			2 1 3`,
			`2`,
		},
		{
			`7
			7 3 6 2 1 5 4
			6 7 2 5 3 1 4`,
			`16`,
		},
		{
			`6
			1 2 3 4 5 6
			6 5 4 3 2 1`,
			`11`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1793D, testCases, 0)
}
