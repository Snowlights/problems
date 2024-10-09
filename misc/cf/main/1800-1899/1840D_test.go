//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1840/D
// https://codeforces.com/problemset/status/1840/problem/D
func TestCF1840D(t *testing.T) {
	t.Log("Current test is [CF1840D]")
	testCases := [][2]string{
		{
			`5
			6
			1 7 7 9 9 9
			6
			5 4 2 1 30 60
			9
			14 19 37 59 1 4 4 98 73
			1
			2
			6
			3 10 1 17 15 11`,
			`0
			2
			13
			0
			1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1840D, testCases, 0)
}
