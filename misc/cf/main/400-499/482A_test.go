//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/482/A
// https://codeforces.com/problemset/status/482/problem/A
func TestCF482A(t *testing.T) {
	t.Log("Current test is [CF482A]")
	testCases := [][2]string{
		{
			`3 2`,
			`1 3 2`,
		},
		{
			`3 1`,
			`1 2 3`,
		},
		{
			`5 2`,
			`1 3 2 4 5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF482A, testCases, 0)
}
