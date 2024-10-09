//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1898/problem/B
// https://codeforces.com/problemset/status/1898/problem/B
func TestCF1898B(t *testing.T) {
	t.Log("Current test is [CF1898B]")
	testCases := [][2]string{
		{
			`4
			3
			1 3 2
			4
			1 2 3 4
			3
			3 2 1
			7
			1 4 4 3 5 7 6`,
			`1
			0
			3
			9
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1898B, testCases, 0)
}
