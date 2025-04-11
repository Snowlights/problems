//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1265/B
// https://codeforces.com/problemset/status/1265/problem/B
func TestCF1265B(t *testing.T) {
	t.Log("Current test is [CF1265B]")
	testCases := [][2]string{
		{
			`3
			6
			4 5 1 3 2 6
			5
			5 3 1 2 4
			4
			1 4 3 2`,
			`101011
			11111
			1001`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1265B, testCases, 0)
}
