//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1878/E
// https://codeforces.com/problemset/status/1878/problem/E
func TestCF1878E(t *testing.T) {
	t.Log("Current test is [CF1878E]")
	testCases := [][2]string{
		{
			`3
			5
			15 14 17 42 34
			3
			1 7
			2 15
			4 5
			5
			7 5 3 1 7
			4
			1 7
			5 7
			2 3
			2 2
			7
			19 20 15 12 21 7 11
			4
			1 15
			4 4
			7 12
			5 7`,
			`2 -1 5 
			1 5 2 2 
			2 6 -1 5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1878E, testCases, 0)
}
