//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1678/problem/B2
// https://codeforces.com/problemset/status/1678/problem/B2
func TestCF1678B2(t *testing.T) {
	t.Log("Current test is [CF1678B2]")
	testCases := [][2]string{
		{
			`5
			10
			1110011000
			8
			11001111
			2
			00
			2
			11
			6
			100110`,
			`3 2
			0 3
			0 1
			0 1
			3 1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1678B2, testCases, 0)
}
