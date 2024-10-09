//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1208/B
// https://codeforces.com/problemset/status/1208/problem/B
func TestCF1208B(t *testing.T) {
	t.Log("Current test is [CF1208B]")
	testCases := [][2]string{
		{
			`3
			1 2 3`,
			`0`,
		},
		{
			`4
			1 1 2 2`,
			`2`,
		},
		{
			`5
			1 4 1 4 9`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1208B, testCases, 0)
}
