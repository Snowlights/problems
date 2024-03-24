//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1706/C
// https://codeforces.com/problemset/status/1706/problem/C
func TestCF1706C(t *testing.T) {
	t.Log("Current test is [CF1706C]")
	testCases := [][2]string{
		{
			`6
			3
			2 1 2
			5
			1 2 1 4 3
			6
			3 1 4 5 5 2
			8
			4 2 1 3 5 3 6 1
			6
			1 10 1 1 10 1
			8
			1 10 11 1 10 11 10 1`,
			`2
			0
			3
			3
			0
			4
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1706C, testCases, 0)
}
