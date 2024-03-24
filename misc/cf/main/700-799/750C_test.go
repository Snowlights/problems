//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/750/C
// https://codeforces.com/problemset/status/750/problem/C
func TestCF750C(t *testing.T) {
	t.Log("Current test is [CF750C]")
	testCases := [][2]string{
		{
			`3
			-7 1
			5 2
			8 2
			`,
			`1907`,
		},
		{
			`2
			57 1
			22 2
			`,
			`Impossible`,
		},
		{
			`1
			-5 1
			`,
			`Infinity`,
		},
		{
			`4
			27 2
			13 1
			-50 1
			8 2
			`,
			`1897`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF750C, testCases, 0)
}
