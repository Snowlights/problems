//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1906/B
// https://codeforces.com/problemset/status/1906/problem/B
func TestCF1906B(t *testing.T) {
	t.Log("Current test is [CF1906B]")
	testCases := [][2]string{
		{
			`2
			4
			0101
			0100
			3
			000
			010
			`,
			`YES
			NO`,
		},
		{
			`5
			7
			0101011
			1111010
			5
			11111
			00000
			4
			1101
			1101
			6
			101010
			100100
			3
			000
			000
			`,
			`NO
			NO
			YES
			YES
			YES
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1906B, testCases, 0)
}
