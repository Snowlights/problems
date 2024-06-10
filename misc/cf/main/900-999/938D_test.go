//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/938/D
// https://codeforces.com/problemset/status/938/problem/D
func TestCF938D(t *testing.T) {
	t.Log("Current test is [CF938D]")
	testCases := [][2]string{
		{
			`4 2
			1 2 4
			2 3 7
			6 20 1 25
			`,
			`6 14 1 25`,
		},
		{
			`3 3
			1 2 1
			2 3 1
			1 3 1
			30 10 20
			`,
			`12 10 12`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF938D, testCases, 0)
}
