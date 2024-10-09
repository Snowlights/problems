//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/452/D
// https://codeforces.com/problemset/status/452/problem/D
func TestCF452D(t *testing.T) {
	t.Log("Current test is [CF452D]")
	testCases := [][2]string{
		{
			`1 1 1 1 5 5 5`,
			`15`,
		},
		{
			`8 4 3 2 10 5 2`,
			`32`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF452D, testCases, 0)
}
