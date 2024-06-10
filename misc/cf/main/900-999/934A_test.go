//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/934/A
// https://codeforces.com/problemset/status/934/problem/A
func TestCF934A(t *testing.T) {
	t.Log("Current test is [CF934A]")
	testCases := [][2]string{
		{
			`2 2
			20 18
			2 14
			`,
			`252`,
		},
		{
			`5 3
			-1 0 1 2 3
			-1 0 1
			`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF934A, testCases, 0)
}
