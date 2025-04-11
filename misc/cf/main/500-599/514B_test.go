//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/514/B
// https://codeforces.com/problemset/status/514/problem/B
func TestCF514B(t *testing.T) {
	t.Log("Current test is [CF514B]")
	testCases := [][2]string{
		{
			`4 0 0
			1 1
			2 2
			2 0
			-1 -1`,
			`2`,
		},
		{
			`2 1 2
			1 1
			1 0`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF514B, testCases, 0)
}
