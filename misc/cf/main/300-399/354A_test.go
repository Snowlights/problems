//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/354/A
// https://codeforces.com/problemset/status/354/problem/A
func TestCF354A(t *testing.T) {
	t.Log("Current test is [CF354A]")
	testCases := [][2]string{
		{
			`3 4 4 19 1
			42 3 99`,
			`576`,
		},
		{
			`4 7 2 3 9
			1 2 3 4`,
			`34`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF354A, testCases, 0)
}
