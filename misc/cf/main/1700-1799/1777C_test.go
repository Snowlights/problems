//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1777/C
// https://codeforces.com/problemset/status/1777/problem/C
func TestCF1777C(t *testing.T) {
	t.Log("Current test is [CF1777C]")
	testCases := [][2]string{
		{
			`3
			2 4
			3 7
			4 2
			3 7 2 9
			5 7
			6 4 3 5 7`,
			`-1
			0
			3
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1777C, testCases, 0)
}
