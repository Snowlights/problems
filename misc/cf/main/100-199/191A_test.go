//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/191/A
// https://codeforces.com/problemset/status/191/problem/A
func TestCF191A(t *testing.T) {
	t.Log("Current test is [CF191A]")
	testCases := [][2]string{
		{
			`3
			abc
			ca
			cba`,
			`6`,
		},
		{
			`4
			vvp
			vvp
			dam
			vvp`,
			`0`,
		},
		{
			`3
			ab
			c
			def`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF191A, testCases, 0)
}
