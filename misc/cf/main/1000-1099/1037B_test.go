//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1037/B
// https://codeforces.com/problemset/status/1037/problem/B
func TestCF1037B(t *testing.T) {
	t.Log("Current test is [CF1037B]")
	testCases := [][2]string{
		{
			`3 8
			6 5 8`,
			`2`,
		},
		{
			`7 20
			21 15 12 11 20 19 12`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1037B, testCases, 0)
}
