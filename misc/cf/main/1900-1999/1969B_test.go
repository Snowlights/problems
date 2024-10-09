//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1969/B
// https://codeforces.com/problemset/status/1969/problem/B
func TestCF1969B(t *testing.T) {
	t.Log("Current test is [CF1969B]")
	testCases := [][2]string{
		{
			`5
			10
			0000
			11000
			101011
			01101001`,
			`2
			0
			9
			5
			11`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1969B, testCases, 0)
}
