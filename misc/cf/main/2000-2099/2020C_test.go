//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2020/C
// https://codeforces.com/problemset/status/2020/problem/C
func TestCF2020C(t *testing.T) {
	t.Log("Current test is [CF2020C]")
	testCases := [][2]string{
		{
			`3
			2 2 2
			4 2 6
			10 2 14`,
			`0
			-1
			12`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2020C, testCases, 0)
}
