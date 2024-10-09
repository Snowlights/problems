//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1903/C
// https://codeforces.com/problemset/status/1903/problem/C
func TestCF1903C(t *testing.T) {
	t.Log("Current test is [CF1903C]")
	testCases := [][2]string{
		{
			`4
			6
			1 -3 7 -6 2 5
			4
			2 9 -5 -3
			8
			-3 -4 2 -5 1 10 17 23
			1
			830`,
			`32
			4
			343
			830
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1903C, testCases, 0)
}
