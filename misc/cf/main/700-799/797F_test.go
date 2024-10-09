//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/797/F
// https://codeforces.com/problemset/status/797/problem/F
func TestCF797F(t *testing.T) {
	t.Log("Current test is [CF797F]")
	testCases := [][2]string{
		{
			`4 5
			6 2 8 9
			3 6
			2 1
			3 6
			4 7
			4 7`,
			`11`,
		},
		{
			`7 2
			10 20 30 40 50 45 35
			-1000000000 10
			1000000000 1`,
			`7000000130`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF797F, testCases, 0)
}
