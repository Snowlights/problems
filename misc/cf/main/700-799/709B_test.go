//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/709/B
// https://codeforces.com/problemset/status/709/problem/B
func TestCF709B(t *testing.T) {
	t.Log("Current test is [CF709B]")
	testCases := [][2]string{
		{
			`3 10
			1 7 12`,
			`7`,
		},
		{
			`2 0
			11 -10`,
			`10`,
		},
		{
			`5 0
			0 0 1000 0 0`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF709B, testCases, 0)
}
