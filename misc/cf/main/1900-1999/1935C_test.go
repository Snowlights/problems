//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1935/C
// https://codeforces.com/problemset/status/1935/problem/C
func TestCF1935C(t *testing.T) {
	t.Log("Current test is [CF1935C]")
	testCases := [][2]string{
		{
			`5
			5 8
			4 3
			1 5
			2 4
			4 3
			2 3
			1 6
			4 10
			3 12
			4 8
			2 1
			2 12
			5 26
			24 7
			8 28
			30 22
			3 8
			17 17
			5 14
			15 3
			1000000000 998244353
			179 239
			228 1337
			993 1007`,
			`3
			1
			2
			1
			0
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1935C, testCases, 0)
}
