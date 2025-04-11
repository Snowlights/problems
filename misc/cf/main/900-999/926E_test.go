//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/926/E
// https://codeforces.com/problemset/status/926/problem/E
func TestCF926E(t *testing.T) {
	t.Log("Current test is [CF926E]")
	testCases := [][2]string{
		{
			`6
			5 2 1 1 2 2`,
			`2
			5 4`,
		},
		{
			`4
			1000000000 1000000000 1000000000 1000000000`,
			`1
			1000000002`,
		},
		{
			`7
			4 10 22 11 12 5 6`,
			`7
			4 10 22 11 12 5 6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF926E, testCases, 0)
}
