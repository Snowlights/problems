//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1811/problem/E
// https://codeforces.com/problemset/status/1811/problem/E
func TestCF1811E(t *testing.T) {
	t.Log("Current test is [CF1811E]")
	testCases := [][2]string{
		{
			`7
			3
			5
			22
			10
			100
			12345
			827264634912`,
			`3
			6
			25
			11
			121
			18937
			2932285320890
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1811E, testCases, 0)
}
