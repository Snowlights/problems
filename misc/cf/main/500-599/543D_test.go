//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/543/D
// https://codeforces.com/problemset/status/543/problem/D
func TestCF543D(t *testing.T) {
	t.Log("Current test is [CF543D]")
	testCases := [][2]string{
		{
			`3
			1 1`,
			`4 3 3`,
		},
		{
			`5
			1 2 3 4`,
			`5 8 9 8 5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF543D, testCases, 0)
}
