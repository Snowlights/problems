//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/862/A
// https://codeforces.com/problemset/status/862/problem/A
func TestCF862A(t *testing.T) {
	t.Log("Current test is [CF862A]")
	testCases := [][2]string{
		{
			`5 3
			0 4 5 6 7`,
			`2`,
		},
		{
			`1 0
			0`,
			`1`,
		},
		{
			`5 0
			1 2 3 4 5`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF862A, testCases, 0)
}
