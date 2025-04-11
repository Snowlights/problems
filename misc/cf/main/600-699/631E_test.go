//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/631/E
// https://codeforces.com/problemset/status/631/problem/E
func TestCF631E(t *testing.T) {
	t.Log("Current test is [CF631E]")
	testCases := [][2]string{
		{
			`4
			4 3 2 5`,
			`39`,
		},
		{
			`5
			1 1 2 7 1`,
			`49`,
		},
		{
			`3
			1 1 2`,
			`9`,
		},
		{
			`7
			0 0 0 -1 0 0 0`,
			`-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF631E, testCases, 0)
}
