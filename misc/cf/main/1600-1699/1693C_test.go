//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1693/C
// https://codeforces.com/problemset/status/1693/problem/C
func TestCF1693C(t *testing.T) {
	t.Log("Current test is [CF1693C]")
	testCases := [][2]string{
		{
			`2 1
			1 2
			`,
			`1`,
		},
		{
			`4 4
			1 2
			1 4
			2 4
			1 4
			`,
			`2`,
		},
		{
			`5 7
			1 2
			2 3
			3 5
			1 4
			4 3
			4 5
			3 1
			`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1693C, testCases, 0)
}
