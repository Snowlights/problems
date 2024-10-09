//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/633/F
// https://codeforces.com/problemset/status/633/problem/F
func TestCF633F(t *testing.T) {
	t.Log("Current test is [CF633F]")
	testCases := [][2]string{
		{
			`9
			1 2 3 4 5 6 7 8 9
			1 2
			1 3
			1 4
			1 5
			1 6
			1 7
			1 8
			1 9`,
			`25`,
		},
		{
			`2
			20 10
			1 2`,
			`30`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF633F, testCases, 0)
}
