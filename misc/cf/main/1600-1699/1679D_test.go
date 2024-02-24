//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1679/D
// https://codeforces.com/problemset/status/1679/problem/D
func TestCF1679D(t *testing.T) {
	t.Log("Current test is [CF1679D]")
	testCases := [][2]string{
		{
			`6 7 4
			1 10 2 3 4 5
			1 2
			1 3
			3 4
			4 5
			5 6
			6 2
			2 5
			`,
			`4
			`,
		},
		{
			`6 7 100
			1 10 2 3 4 5
			1 2
			1 3
			3 4
			4 5
			5 6
			6 2
			2 5
			`,
			`10`,
		},
		{
			`2 1 5
			1 1
			1 2
			`,
			`-1`,
		},
		{
			`1 0 1
			1000000000
			`,
			`1000000000`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1679D, testCases, 0)
}
