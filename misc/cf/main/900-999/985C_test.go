//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/985/C
// https://codeforces.com/problemset/status/985/problem/C
func TestCF985C(t *testing.T) {
	t.Log("Current test is [CF985C]")
	testCases := [][2]string{
		{
			`4 2 1
			2 2 1 2 3 2 2 3`,
			`7`,
		},
		{
			`2 1 0
			10 10`,
			`20`,
		},
		{
			`1 2 1
			5 2`,
			`2`,
		},
		{
			`3 2 1
			1 2 3 4 5 6`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF985C, testCases, 0)
}
