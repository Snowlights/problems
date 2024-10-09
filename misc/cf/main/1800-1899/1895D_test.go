//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1895/D
// https://codeforces.com/problemset/status/1895/problem/D
func TestCF1895D(t *testing.T) {
	t.Log("Current test is [CF1895D]")
	testCases := [][2]string{
		{
			`4
			2 1 2`,
			`0 2 3 1`,
		},
		{
			`6
			1 6 1 4 1`,
			`2 3 5 4 0 1 `,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1895D, testCases, 0)
}
