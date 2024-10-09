//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/696/B
// https://codeforces.com/problemset/status/696/problem/B
func TestCF696B(t *testing.T) {
	t.Log("Current test is [CF696B]")
	testCases := [][2]string{
		{
			`7
			1 2 1 1 4 4
			`,
			`1.0 4.0 5.0 3.5 4.5 5.0 5.0`,
		},
		{
			`12
			1 1 2 2 4 4 3 3 1 10 8
			`,
			`1.0 5.0 5.5 6.5 7.5 8.0 8.0 7.0 7.5 6.5 7.5 8.0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF696B, testCases, 0)
}
