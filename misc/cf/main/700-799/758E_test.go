//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/758/E
// https://codeforces.com/problemset/status/758/problem/E
func TestCF758E(t *testing.T) {
	t.Log("Current test is [CF758E]")
	testCases := [][2]string{
		{
			`3
			1 3 5 7
			3 2 4 3`,
			`3
			1 3 5 7
			3 2 4 3`,
		},
		{
			`4
			1 3 2 3
			3 4 5 1
			3 2 3 3`,
			`-1`,
		},
		{
			`5
			1 2 2 4
			2 4 1 9
			4 5 5 6
			4 3 4 8`,
			`5
			1 2 2 4
			2 4 1 9
			4 5 1 2
			4 3 2 6`,
		},
		{
			`7
			1 2 5 2
			2 3 4 3
			1 4 3 7
			4 5 4 1
			4 6 3 2
			6 7 1 6`,
			`7
			1 2 5 2
			2 3 2 1
			1 4 3 7
			4 5 3 0
			4 6 3 2
			6 7 1 6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF758E, testCases, 0)
}
