//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1789/C
// https://codeforces.com/problemset/status/1789/problem/C
func TestCF1789C(t *testing.T) {
	t.Log("Current test is [CF1789C]")
	testCases := [][2]string{
		{
			`3
			3 2
			1 2 3
			1 4
			2 5
			1 1
			1
			1 1
			10 10
			4 6 9 12 16 20 2 10 19 7
			1 3
			5 4
			2 17
			2 18
			6 11
			7 1
			8 17
			5 5
			5 5
			2 2`,
			`13
			1
			705`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1789C, testCases, 0)
}
