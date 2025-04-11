//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1990/C
// https://codeforces.com/problemset/status/1990/problem/C
func TestCF1990C(t *testing.T) {
	t.Log("Current test is [CF1990C]")
	testCases := [][2]string{
		{
			`4
			1
			1
			3
			2 2 3
			4
			2 1 1 2
			4
			4 4 4 4`,
			`1
			13
			9
			40`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1990C, testCases, 0)
}
