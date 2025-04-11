//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/579/A
// https://codeforces.com/problemset/status/579/problem/A
func TestCF579A(t *testing.T) {
	t.Log("Current test is [CF579A]")
	testCases := [][2]string{
		{
			`5`,
			`2`,
		},
		{
			`8`,
			`1`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF579A, testCases, 0)
}
