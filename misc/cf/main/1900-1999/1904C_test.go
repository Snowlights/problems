//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1904/C
// https://codeforces.com/problemset/status/1904/problem/C
func TestCF1904C(t *testing.T) {
	t.Log("Current test is [CF1904C]")
	testCases := [][2]string{
		{
			`4
			5 2
			3 9 7 15 1
			4 3
			7 4 15 12
			6 2
			42 47 50 54 62 79
			2 1
			500000000000000000 1000000000000000000`,
			`1
			0
			3
			500000000000000000`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1904C, testCases, 0)
}
