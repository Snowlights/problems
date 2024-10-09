//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1997/D
// https://codeforces.com/problemset/status/1997/problem/D
func TestCF1997D(t *testing.T) {
	t.Log("Current test is [CF1997D]")
	testCases := [][2]string{
		{
			`3
			4
			0 1 0 2
			1 1 3
			2
			3 0
			1
			5
			2 5 3 9 6
			3 1 5 2`,
			`1
			3
			6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1997D, testCases, 0)
}
