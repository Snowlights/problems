//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/212/E
// https://codeforces.com/problemset/status/212/problem/E
func TestCF212E(t *testing.T) {
	t.Log("Current test is [CF212E]")
	testCases := [][2]string{
		{
			`5
			1 2
			2 3
			3 4
			4 5`,
			`3
			1 3
			2 2
			3 1`,
		},
		{
			`10
			1 2
			2 3
			3 4
			5 6
			6 7
			7 4
			8 9
			9 10
			10 4`,
			`6
			1 8
			2 7
			3 6
			6 3
			7 2
			8 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF212E, testCases, 0)
}
