//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1227/G
// https://codeforces.com/problemset/status/1227/problem/G
func TestCF1227G(t *testing.T) {
	t.Log("Current test is [CF1227G]")
	testCases := [][2]string{
		{
			`5
			5 5 5 5 5`,
			`6
			11111
			01111
			10111
			11011
			11101
			11110`,
		},
		{
			`5
			5 1 1 1 1`,
			`5
			11000
			10000
			10100
			10010
			10001`,
		},
		{
			`5
			4 1 5 3 4`,
			`5
			11111
			10111
			10101
			00111
			10100`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1227G, testCases, 0)
}
