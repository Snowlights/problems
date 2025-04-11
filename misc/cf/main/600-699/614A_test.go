//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/614/A
// https://codeforces.com/problemset/status/614/problem/A
func TestCF614A(t *testing.T) {
	t.Log("Current test is [CF614A]")
	testCases := [][2]string{
		{
			`1 10 2`,
			`1 2 4 8`,
		},
		{
			`2 4 5`,
			`-1`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF614A, testCases, 0)
}
