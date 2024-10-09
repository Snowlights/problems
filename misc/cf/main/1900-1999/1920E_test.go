//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1920/E
// https://codeforces.com/problemset/status/1920/problem/E
func TestCF1920E(t *testing.T) {
	t.Log("Current test is [CF1920E]")
	testCases := [][2]string{
		{
			`6
			1 1
			3 2
			4 2
			5 4
			6 2
			2450 2391`,
			`1
			3
			5
			12
			9
			259280854`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1920E, testCases, 0)
}
