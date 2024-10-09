//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1950/F
// https://codeforces.com/problemset/status/1950/problem/F
func TestCF1950F(t *testing.T) {
	t.Log("Current test is [CF1950F]")
	testCases := [][2]string{
		{
			`10
			2 1 3
			0 0 1
			0 1 1
			1 0 2
			1 1 3
			3 1 4
			8 17 9
			24 36 48
			1 0 0
			0 3 1`,
			`2
			0
			1
			1
			-1
			3
			6
			-1
			-1
			3
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1950F, testCases, 0)
}
