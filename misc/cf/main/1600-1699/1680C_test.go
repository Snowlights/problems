//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1680/problem/C
// https://codeforces.com/problemset/status/1680/problem/C
func TestCF1680C(t *testing.T) {
	t.Log("Current test is [CF1680C]")
	testCases := [][2]string{
		{
			`5
			101110110
			1001001001001
			0000111111
			00000
			1111`,
			`1
			3
			0
			0
			0
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1680C, testCases, 0)
}
