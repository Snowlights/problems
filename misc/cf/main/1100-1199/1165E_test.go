//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1165/E
// https://codeforces.com/problemset/status/1165/problem/E
func TestCF1165E(t *testing.T) {
	t.Log("Current test is [CF1165E]")
	testCases := [][2]string{
		{
			`5
			1 8 7 2 4
			9 7 2 9 3
			`,
			`646`,
		},
		{
			`1
			1000000
			1000000
			`,
			`757402647`,
		},
		{
			`2
			1 3
			4 2
			`,
			`20`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1165E, testCases, 0)
}
