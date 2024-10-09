//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/383/A
// https://codeforces.com/problemset/status/383/problem/A
func TestCF383A(t *testing.T) {
	t.Log("Current test is [CF383A]")
	testCases := [][2]string{
		{
			`4
			0 0 1 0
			`,
			`1`,
		},
		{
			`5
			1 0 1 0 1
			`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF383A, testCases, 0)
}
