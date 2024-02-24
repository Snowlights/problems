//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/404/problem/D
// https://codeforces.com/problemset/status/404/problem/D
func TestCF404D(t *testing.T) {
	t.Log("Current test is [CF404D]")
	testCases := [][2]string{
		{
			`?01???`,
			`4`,
		},
		{
			`?`,
			`2`,
		},
		{
			`**12`,
			`0`,
		},
		{
			`1`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF404D, testCases, 0)
}
