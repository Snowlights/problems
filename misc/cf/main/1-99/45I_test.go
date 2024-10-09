//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/45/I
// https://codeforces.com/problemset/status/45/problem/I
func TestCF45I(t *testing.T) {
	t.Log("Current test is [CF45I]")
	testCases := [][2]string{
		{
			`5
			1 2 -3 3 3
			`,
			`3 1 2 3`,
		},
		{
			`13
			100 100 100 100 100 100 100 100 100 100 100 100 100
		`,
			`100 100 100 100 100 100 100 100 100 100 100 100 100`,
		},
		{
			`4
			-2 -2 -2 -2
			`,
			`-2 -2 -2 -2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF45I, testCases, 0)
}
