//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1765/K
// https://codeforces.com/problemset/status/1765/problem/K
func TestCF1765K(t *testing.T) {
	t.Log("Current test is [CF1765K]")
	testCases := [][2]string{
		{
			`2
			1 2
			3 4`,
			`8`,
		},
		{
			`3
			10 10 10
			10 0 10
			10 10 10`,
			`80`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1765K, testCases, 0)
}
