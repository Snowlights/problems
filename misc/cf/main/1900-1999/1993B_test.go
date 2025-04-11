//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1993/B
// https://codeforces.com/problemset/status/1993/problem/B
func TestCF1993B(t *testing.T) {
	t.Log("Current test is [CF1993B]")
	testCases := [][2]string{
		{
			`7
			5
			1 3 5 7 9
			4
			4 4 4 4
			3
			2 3 4
			4
			3 2 2 8
			6
			4 3 6 1 2 1
			6
			3 6 1 2 1 2
			5
			999999996 999999997 999999998 999999999 1000000000`,
			`0
			0
			2
			4
			3
			3
			3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1993B, testCases, 0)
}
