//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1167/E
// https://codeforces.com/problemset/status/1167/problem/E
func TestCF1167E(t *testing.T) {
	t.Log("Current test is [CF1167E]")
	testCases := [][2]string{
		{
			`3 3
			2 3 1`,
			`4
			`,
		},
		{
			`7 4
			1 3 1 2 2 4 3`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1167E, testCases, 0)
}
