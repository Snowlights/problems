//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1513/B
// https://codeforces.com/problemset/status/1513/problem/B
func TestCF1513B(t *testing.T) {
	t.Log("Current test is [CF1513B]")
	testCases := [][2]string{
		{
			`4
			3
			1 1 1
			5
			1 2 3 4 5
			5
			0 2 0 3 0
			4
			1 3 5 1`,
			`6
			0
			36
			4
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1513B, testCases, 0)
}
