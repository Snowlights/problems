//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1237/F
// https://codeforces.com/problemset/status/1237/problem/F
func TestCF1237F(t *testing.T) {
	t.Log("Current test is [CF1237F]")
	testCases := [][2]string{
		{
			`5 7 2
			3 1 3 2
			4 4 4 5`,
			`8`,
		},
		{
			`5 4 2
			1 2 2 2
			4 3 4 4`,
			`1`,
		},
		{
			`23 42 0`,
			`102848351`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1237F, testCases, 0)
}
