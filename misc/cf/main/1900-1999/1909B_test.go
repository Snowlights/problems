//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1909/B
// https://codeforces.com/problemset/status/1909/problem/B
func TestCF1909B(t *testing.T) {
	t.Log("Current test is [CF1909B]")
	testCases := [][2]string{
		{
			`5
			4
			8 15 22 30
			5
			60 90 98 120 308
			6
			328 769 541 986 215 734
			5
			1000 2000 7000 11000 16000
			2
			2 1`,
			`7
			30
			3
			5000
			1000000000000000000`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1909B, testCases, 0)
}
