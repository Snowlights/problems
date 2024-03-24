//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/702/C
// https://codeforces.com/problemset/status/702/problem/C
func TestCF702C(t *testing.T) {
	t.Log("Current test is [CF702C]")
	testCases := [][2]string{
		{
			`3 2
			-2 2 4
			-3 0
			`,
			`4`,
		},
		{
			`5 3
			1 5 10 14 17
			4 11 15
			`,
			`3`,
		},
		{
			`1 1
			-1000000000
			1000000000`,
			`2000000000`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF702C, testCases, 0)
}
