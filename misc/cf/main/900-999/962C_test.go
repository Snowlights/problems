//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/962/problem/C
// https://codeforces.com/problemset/status/962/problem/C
func TestCF962C(t *testing.T) {
	t.Log("Current test is [CF962C]")
	testCases := [][2]string{
		{
			`8314`,
			`2`,
		},
		{
			`625`,
			`0`,
		},
		{
			`333`,
			`-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF962C, testCases, 0)
}
