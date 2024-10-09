//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1922/C
// https://codeforces.com/problemset/status/1922/problem/C
func TestCF1922C(t *testing.T) {
	t.Log("Current test is [CF1922C]")
	testCases := [][2]string{
		{
			`1
			5
			0 8 12 15 20
			5
			1 4
			1 5
			3 4
			3 2
			5 1`,
			`3
			8
			1
			4
			14
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1922C, testCases, 0)
}
