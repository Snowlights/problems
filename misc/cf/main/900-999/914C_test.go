//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/914/C
// https://codeforces.com/problemset/status/914/problem/C
func TestCF914C(t *testing.T) {
	t.Log("Current test is [CF914C]")
	testCases := [][2]string{
		{
			`110
			2`,
			`3`,
		},
		{
			`111111011
			2`,
			`169`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF914C, testCases, 0)
}
