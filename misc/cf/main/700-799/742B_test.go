//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/742/B
// https://codeforces.com/problemset/status/742/problem/B
func TestCF742B(t *testing.T) {
	t.Log("Current test is [CF742B]")
	testCases := [][2]string{
		{
			`2 3
1 2`,
			`1`,
		},
		{
			`6 1
5 1 2 3 4 1`,
			`2`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF742B, testCases, 0)
}
