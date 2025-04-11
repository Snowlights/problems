//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2044/H
// https://codeforces.com/problemset/status/2044/problem/H
func TestCF2044H(t *testing.T) {
	t.Log("Current test is [CF2044H]")
	testCases := [][2]string{
		{
			`2
			4 3
			1 5 2 4
			4 9 5 3
			4 5 2 3
			1 5 5 2
			1 1 4 4
			2 2 3 3
			1 2 4 3
			3 3
			1 2 3
			4 5 6
			7 8 9
			1 1 1 3
			1 3 3 3
			2 2 2 2`,
			`500 42 168 
			14 42 5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2044H, testCases, 0)
}
