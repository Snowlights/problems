//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1648/A
// https://codeforces.com/problemset/status/1648/problem/A
func TestCF1648A(t *testing.T) {
	t.Log("Current test is [CF1648A]")
	testCases := [][2]string{
		{
			`2 3
			1 2 3
			3 2 1
			`,
			`7`,
		},
		{
			`3 4
			1 1 2 2
			2 1 1 2
			2 2 1 1
			`,
			`76`,
		},
		{
			`4 4
			1 1 2 3
			2 1 1 2
			3 1 2 1
			1 1 2 1
			`,
			`129`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1648A, testCases, 0)
}
