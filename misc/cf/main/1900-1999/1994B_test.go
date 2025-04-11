//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1994/B
// https://codeforces.com/problemset/status/1994/problem/B
func TestCF1994B(t *testing.T) {
	t.Log("Current test is [CF1994B]")
	testCases := [][2]string{
		{
			`6
			1
			0
			1
			7
			0110100
			0110100
			9
			100101010
			101111110
			4
			0011
			1011
			4
			0100
			0001
			8
			10110111
			01100000`,
			`NO
			YES
			YES
			NO
			YES
			YES`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1994B, testCases, 0)
}
