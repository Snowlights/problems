//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1137/B
// https://codeforces.com/problemset/status/1137/problem/B
func TestCF1137B(t *testing.T) {
	t.Log("Current test is [CF1137B]")
	testCases := [][2]string{
		{
			`101101
			110`,
			``,
		},
		{
			`10010110
			100011
			`,
			`01100011`,
		},
		{
			`10
			11100
			`,
			`01`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1137B, testCases, 0)
}
