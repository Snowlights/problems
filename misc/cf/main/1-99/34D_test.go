//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/34/D
// https://codeforces.com/problemset/status/34/problem/D
func TestCF34D(t *testing.T) {
	t.Log("Current test is [CF34D]")
	testCases := [][2]string{
		{
			`3 2 3
			2 2`,
			`2 3`,
		},
		{
			`6 2 4
			6 1 2 4 2`,
			`6 4 1 4 2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF34D, testCases, 0)
}
