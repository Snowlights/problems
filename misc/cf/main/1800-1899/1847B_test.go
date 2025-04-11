//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1847/B
// https://codeforces.com/problemset/status/1847/problem/B
func TestCF1847B(t *testing.T) {
	t.Log("Current test is [CF1847B]")
	testCases := [][2]string{
		{
			`3
			3
			1 2 3
			5
			2 3 1 5 2
			4
			5 7 12 6`,
			`1
			2
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1847B, testCases, 0)
}
