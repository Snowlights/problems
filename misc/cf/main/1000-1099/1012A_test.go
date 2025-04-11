//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1012/A
// https://codeforces.com/problemset/status/1012/problem/A
func TestCF1012A(t *testing.T) {
	t.Log("Current test is [CF1012A]")
	testCases := [][2]string{
		{
			`4
			4 1 3 2 3 2 1 3`,
			`1`,
		},
		{
			`3
			5 8 5 5 7 5`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1012A, testCases, 0)
}
