//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/698/A
// https://codeforces.com/problemset/status/698/problem/A
func TestCF698A(t *testing.T) {
	t.Log("Current test is [CF698A]")
	testCases := [][2]string{
		{
			`4
			1 3 2 0`,
			`2`,
		},
		{
			`7
			1 3 3 2 1 2 3`,
			`0`,
		},
		{
			`2
			2 2`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF698A, testCases, 0)
}
