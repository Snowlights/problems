//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2047/A
// https://codeforces.com/problemset/status/2047/problem/A
func TestCF2047A(t *testing.T) {
	t.Log("Current test is [CF2047A]")
	testCases := [][2]string{
		{
			`5
			1
			1
			2
			1 8
			5
			1 3 2 1 2
			7
			1 2 1 10 2 7 2
			14
			1 10 10 100 1 1 10 1 10 2 10 2 10 1`,
			`1
			2
			2
			2
			3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2047A, testCases, 0)
}
