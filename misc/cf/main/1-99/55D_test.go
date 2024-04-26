//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/55/D
// https://codeforces.com/problemset/status/55/problem/D
func TestCF55D(t *testing.T) {
	t.Log("Current test is [CF55D]")
	testCases := [][2]string{
		{
			`1
			1 9
			`,
			`9`,
		},
		{
			`1
			12 15
			`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF55D, testCases, 0)
}
