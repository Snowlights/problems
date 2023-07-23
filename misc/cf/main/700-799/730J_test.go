//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/730/J
// https://codeforces.com/problemset/status/730/problem/J
func TestCF730J(t *testing.T) {
	t.Log("Current test is [CF730J]")
	testCases := [][2]string{
		{
			`4
			3 3 4 3
			4 7 6 5`,
			`2 6`,
		},
		{
			`2
			1 1
			100 100`,
			`1 1`,
		},
		{
			`5
			10 30 5 6 24
			10 41 7 8 24`,
			`3 11`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF730J, testCases, 0)
}
