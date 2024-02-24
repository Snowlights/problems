//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/525/C
// https://codeforces.com/problemset/status/525/problem/C
func TestCF525C(t *testing.T) {
	t.Log("Current test is [CF525C]")
	testCases := [][2]string{
		{
			`4
			2 4 4 2
			`,
			`8`,
		},
		{
			`4
			2 2 3 5
			`,
			`0`,
		},
		{
			`4
			100003 100004 100005 100006
			`,
			`10000800015`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF525C, testCases, 0)
}
