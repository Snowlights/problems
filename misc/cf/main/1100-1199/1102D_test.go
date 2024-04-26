//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1102/D
// https://codeforces.com/problemset/status/1102/problem/D
func TestCF1102D(t *testing.T) {
	t.Log("Current test is [CF1102D]")
	testCases := [][2]string{
		{
			`3
			121
			`,
			`021`,
		},
		{
			`6
			000000
			`,
			`001122
			`,
		},
		{
			`6
			211200
			`,
			`211200
			`,
		},
		{
			`6
			120110
			`,
			`120120
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1102D, testCases, 0)
}
