//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/976/C
// https://codeforces.com/problemset/status/976/problem/C
func TestCF976C(t *testing.T) {
	t.Log("Current test is [CF976C]")
	testCases := [][2]string{
		{
			`5
			1 10
			2 9
			3 9
			2 3
			2 9`,
			`2 1`,
		},
		{
			`3
			1 5
			2 6
			6 20`,
			`-1 -1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF976C, testCases, 0)
}
