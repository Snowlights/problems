//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2026/B
// https://codeforces.com/problemset/status/2026/problem/B
func TestCF2026B(t *testing.T) {
	t.Log("Current test is [CF2026B]")
	testCases := [][2]string{
		{
			`4
			2
			1 2
			1
			7
			3
			2 4 9
			5
			1 5 8 10 13`,
			`1
			1
			2
			3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2026B, testCases, 0)
}
