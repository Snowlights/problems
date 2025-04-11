//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1822/G2
// https://codeforces.com/problemset/status/1822/problem/G2
func TestCF1822G2(t *testing.T) {
	t.Log("Current test is [CF1822G2]")
	testCases := [][2]string{
		{
			`7
			5
			1 7 7 2 7
			3
			6 2 18
			9
			1 2 3 4 5 6 7 8 9
			4
			1000 993 986 179
			7
			1 10 100 1000 10000 100000 1000000
			8
			1 1 2 2 4 4 8 8
			9
			1 1 1 2 2 2 4 4 4`,
			`6
			1
			3
			0
			9
			16
			45`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1822G2, testCases, 0)
}
