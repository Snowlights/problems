//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2008/F
// https://codeforces.com/problemset/status/2008/problem/F
func TestCF2008F(t *testing.T) {
	t.Log("Current test is [CF2008F]")
	testCases := [][2]string{
		{
			`3
			3
			3 2 3
			4
			2 2 2 4
			5
			1 2 3 4 5`,
			`7
			6
			500000012`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2008F, testCases, 0)
}
