//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/816/B
// https://codeforces.com/problemset/status/816/problem/B
func TestCF816B(t *testing.T) {
	t.Log("Current test is [CF816B]")
	testCases := [][2]string{
		{
			`3 2 4
			91 94
			92 97
			97 99
			92 94
			93 97
			95 96
			90 100
			`,
			`3
			3
			0
			4
			`,
		},
		{
			`2 1 1
			1 1
			200000 200000
			90 100
			`,
			`0
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF816B, testCases, 0)
}
