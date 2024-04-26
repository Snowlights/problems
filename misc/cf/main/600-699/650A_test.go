//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/650/A
// https://codeforces.com/problemset/status/650/problem/A
func TestCF650A(t *testing.T) {
	t.Log("Current test is [CF650A]")
	testCases := [][2]string{
		{
			`3
			1 1
			7 5
			1 5`,
			`2`,
		},
		{
			`6
			0 0
			0 1
			0 2
			-1 1
			0 1
			1 1`,
			`11`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF650A, testCases, 0)
}
