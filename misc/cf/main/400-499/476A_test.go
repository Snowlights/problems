//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/476/A
// https://codeforces.com/problemset/status/476/problem/A
func TestCF476A(t *testing.T) {
	t.Log("Current test is [CF476A]")
	testCases := [][2]string{
		{
			`10 2`,
			`6`,
		},
		{
			`3 5`,
			`-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF476A, testCases, 0)
}
