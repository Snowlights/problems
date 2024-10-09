//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1873/F
// https://codeforces.com/problemset/status/1873/problem/F
func TestCF1873F(t *testing.T) {
	t.Log("Current test is [CF1873F]")
	testCases := [][2]string{
		{
			`5
			5 12
			3 2 4 1 8
			4 4 2 4 1
			4 8
			5 4 1 2
			6 2 3 1
			3 12
			7 9 10
			2 2 4
			1 10
			11
			1
			7 10
			2 6 3 1 5 10 6
			72 24 24 12 4 4 2`,
			`3
			2
			1
			0
			3
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1873F, testCases, 0)
}
