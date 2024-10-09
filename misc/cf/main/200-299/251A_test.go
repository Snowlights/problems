//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/251/A
// https://codeforces.com/problemset/status/251/problem/A
func TestCF251A(t *testing.T) {
	t.Log("Current test is [CF251A]")
	testCases := [][2]string{
		{
			`4 3
			1 2 3 4`,
			`4`,
		},
		{
			`4 2
			-3 -2 -1 0`,
			`2`,
		},
		{
			`5 19
			1 10 20 30 50`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF251A, testCases, 0)
}
