//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/204/A
// https://codeforces.com/problemset/status/204/problem/A
func TestCF204A(t *testing.T) {
	t.Log("Current test is [CF204A]")
	testCases := [][2]string{
		{
			`2 47`,
			`12`,
		},
		{
			`47 1024`,
			`98`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF204A, testCases, 0)
}
