//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1881/E
// https://codeforces.com/problemset/status/1881/problem/E
func TestCF1881E(t *testing.T) {
	t.Log("Current test is [CF1881E]")
	testCases := [][2]string{
		{
			`7
			7
			3 3 4 5 2 6 1
			4
			5 6 3 2
			6
			3 4 1 6 7 7
			3
			1 4 3
			5
			1 2 3 4 5
			5
			1 2 3 1 2
			5
			4 5 5 1 5`,
			`0
			4
			1
			1
			2
			1
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1881E, testCases, 0)
}
