//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/238/C
// https://codeforces.com/problemset/status/238/problem/C
func TestCF238C(t *testing.T) {
	t.Log("Current test is [CF238C]")
	testCases := [][2]string{
		{
			`4
			2 1
			3 1
			4 1`,
			`1`,
		},
		{
			`5
			2 1
			2 3
			4 3
			4 5`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF238C, testCases, 0)
}
