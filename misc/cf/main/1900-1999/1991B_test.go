//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1991/B
// https://codeforces.com/problemset/status/1991/problem/B
func TestCF1991B(t *testing.T) {
	t.Log("Current test is [CF1991B]")
	testCases := [][2]string{
		{
			`4
			2
			1
			3
			2 0
			4
			1 2 3
			5
			3 5 4 2`,
			`5 3
			3 2 1
			-1
			3 7 5 6 3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1991B, testCases, 0)
}
