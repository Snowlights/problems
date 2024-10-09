//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1765/D
// https://codeforces.com/problemset/status/1765/problem/D
func TestCF1765D(t *testing.T) {
	t.Log("Current test is [CF1765D]")
	testCases := [][2]string{
		{
			`5 6
			1 2 3 4 5`,
			`16`,
		},
		{
			`5 5
			1 2 3 4 5`,
			`17`,
		},
		{
			`4 3
			1 3 2 3`,
			`12`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1765D, testCases, 0)
}
