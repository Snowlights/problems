//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1965/B
// https://codeforces.com/problemset/status/1965/problem/B
func TestCF1965B(t *testing.T) {
	t.Log("Current test is [CF1965B]")
	testCases := [][2]string{
		{
			`5
			2 2
			6 1
			8 8
			9 3
			10 7`,
			`1
			1
			5
			2 3 4 5 6
			7
			1 1 1 1 1 1 1
			4
			7 1 4 1
			4
			1 2 8 3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1965B, testCases, 0)
}
