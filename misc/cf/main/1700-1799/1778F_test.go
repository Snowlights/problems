//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1778/F
// https://codeforces.com/problemset/status/1778/problem/F
func TestCF1778F(t *testing.T) {
	t.Log("Current test is [CF1778F]")
	testCases := [][2]string{
		{
			`2
			5 2
			24 12 24 6 12
			1 2
			1 3
			2 4
			2 5
			5 3
			24 12 24 6 12
			1 2
			1 3
			2 4
			2 5`,
			`288
			576
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1778F, testCases, 0)
}
