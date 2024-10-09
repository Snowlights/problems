//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1883/F
// https://codeforces.com/problemset/status/1883/problem/F
func TestCF1883F(t *testing.T) {
	t.Log("Current test is [CF1883F]")
	testCases := [][2]string{
		{
			`6
			1
			1
			2
			1 1
			3
			1 2 1
			4
			2 3 2 1
			5
			4 5 4 5 4
			10
			1 7 7 2 3 4 3 2 1 100`,
			`1
			1
			4
			7
			4
			28
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1883F, testCases, 0)
}
