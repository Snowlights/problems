//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1304/C
// https://codeforces.com/problemset/status/1304/problem/C
func TestCF1304C(t *testing.T) {
	t.Log("Current test is [CF1304C]")
	testCases := [][2]string{
		{
			`4
			3 0
			5 1 2
			7 3 5
			10 -1 0
			2 12
			5 7 10
			10 16 20
			3 -100
			100 0 0
			100 -50 50
			200 100 100
			1 100
			99 -100 0
			`,
			`YES
			NO
			YES
			NO
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1304C, testCases, 0)
}
