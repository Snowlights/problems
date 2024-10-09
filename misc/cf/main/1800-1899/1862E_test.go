//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1862/E
// https://codeforces.com/problemset/status/1862/problem/E
func TestCF1862E(t *testing.T) {
	t.Log("Current test is [CF1862E]")
	testCases := [][2]string{
		{
			`6
			5 2 2
			3 2 5 4 6
			4 3 2
			1 1 1 1
			6 6 6
			-82 45 1 -77 39 11
			5 2 2
			3 2 5 4 8
			2 1 1
			-1 2
			6 3 2
			-8 8 -2 -1 9 0`,
			`2
			0
			60
			3
			0
			7
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1862E, testCases, 0)
}
