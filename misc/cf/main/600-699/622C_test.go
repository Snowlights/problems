//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/622/C
// https://codeforces.com/problemset/status/622/problem/C
func TestCF622C(t *testing.T) {
	t.Log("Current test is [CF622C]")
	testCases := [][2]string{
		{
			`6 4
			1 2 1 1 3 5
			1 4 1
			2 6 2
			3 4 1
			3 4 2`,
			`2
			6
			-1
			4`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF622C, testCases, 0)
}
