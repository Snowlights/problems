//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1626/C
// https://codeforces.com/problemset/status/1626/problem/C
func TestCF1626C(t *testing.T) {
	t.Log("Current test is [CF1626C]")
	testCases := [][2]string{
		{
			`3
			1
			6
			4
			2
			4 5
			2 2
			3
			5 7 9
			2 1 2`,
			`10
			6
			7
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1626C, testCases, 0)
}
