//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1029/C
// https://codeforces.com/problemset/status/1029/problem/C
func TestCF1029C(t *testing.T) {
	t.Log("Current test is [CF1029C]")
	testCases := [][2]string{
		{
			`4
			1 3
			2 6
			0 4
			3 3
			`,
			`1`,
		},
		{
			`5
			2 6
			1 3
			0 4
			1 20
			0 4
			`,
			`2`,
		},
		{
			`3
			4 5
			1 2
			9 20
			`,
			`0`,
		},
		{
			`2
			3 10
			1 5
			`,
			`7`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1029C, testCases, 0)
}
