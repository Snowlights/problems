//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1701/C
// https://codeforces.com/problemset/status/1701/problem/C
func TestCF1701C(t *testing.T) {
	t.Log("Current test is [CF1701C]")
	testCases := [][2]string{
		{
			`4
			2 4
			1 2 1 2
			2 4
			1 1 1 1
			5 5
			5 1 3 2 4
			1 1
			1`,
			`2
			3
			1
			1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1701C, testCases, 0)
}
