//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1148/B
// https://codeforces.com/problemset/status/1148/problem/B
func TestCF1148B(t *testing.T) {
	t.Log("Current test is [CF1148B]")
	testCases := [][2]string{
		{
			`4 5 1 1 2
			1 3 5 7
			1 2 3 9 10`,
			`11`,
		}, {
			`2 2 4 4 2
			1 10
			10 20`,
			`-1`,
		},
		{
			`4 3 2 3 1
			1 999999998 999999999 1000000000
			3 4 1000000000`,
			`1000000003`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1148B, testCases, 0)
}
