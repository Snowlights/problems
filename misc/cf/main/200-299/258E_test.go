//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/258/E
// https://codeforces.com/problemset/status/258/problem/E
func TestCF258E(t *testing.T) {
	t.Log("Current test is [CF258E]")
	testCases := [][2]string{
		{
			`5 1
			1 2
			1 3
			3 5
			3 4
			2 3`,
			`0 3 3 3 3`,
		},
		{
			`11 3
			1 2
			2 3
			2 4
			1 5
			5 6
			5 7
			5 8
			6 9
			8 10
			8 11
			2 9
			3 6
			2 8`,
			`0 6 7 6 0 2 0 5 4 5 5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF258E, testCases, 0)
}
