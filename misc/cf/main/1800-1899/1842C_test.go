//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1842/C
// https://codeforces.com/problemset/status/1842/problem/C
func TestCF1842C(t *testing.T) {
	t.Log("Current test is [CF1842C]")
	testCases := [][2]string{
		{
			`2
			5
			1 2 2 3 3
			4
			1 2 1 2`,
			`4
			3`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1842C, testCases, 0)
}
