//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2048/C
// https://codeforces.com/problemset/status/2048/problem/C
func TestCF2048C(t *testing.T) {
	t.Log("Current test is [CF2048C]")
	testCases := [][2]string{
		{
			`5
			111
			1000
			10111
			11101
			1100010001101`,
			`2 2 1 3
			1 3 1 4
			1 5 1 4
			3 4 1 5
			1 13 1 11`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2048C, testCases, 0)
}
