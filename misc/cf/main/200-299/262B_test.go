//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/262/B
// https://codeforces.com/problemset/status/262/problem/B
func TestCF262B(t *testing.T) {
	t.Log("Current test is [CF262B]")
	testCases := [][2]string{
		{
			`3 2
			-1 -1 1`,
			`3`,
		},
		{
			`3 1
			-1 -1 1`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF262B, testCases, 0)
}
