//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1215/B
// https://codeforces.com/problemset/status/1215/problem/B
func TestCF1215B(t *testing.T) {
	t.Log("Current test is [CF1215B]")
	testCases := [][2]string{
		{
			`5
			5 -3 3 -1 1`,
			`8 7`,
		},
		{
			`10
			4 2 -4 3 1 2 -4 3 2 3`,
			`28 27`,
		},
		{
			`5
			-1 -2 -3 -4 -5`,
			`9 6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1215B, testCases, 0)
}
