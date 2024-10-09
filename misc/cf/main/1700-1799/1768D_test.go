//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1768/D
// https://codeforces.com/problemset/status/1768/problem/D
func TestCF1768D(t *testing.T) {
	t.Log("Current test is [CF1768D]")
	testCases := [][2]string{
		{
			`4
			2
			2 1
			2
			1 2
			4
			3 4 1 2
			4
			2 4 3 1`,
			`0
			1
			3
			1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1768D, testCases, 0)
}
