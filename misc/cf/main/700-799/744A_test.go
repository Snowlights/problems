//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/744/A
// https://codeforces.com/problemset/status/744/problem/A
func TestCF744A(t *testing.T) {
	t.Log("Current test is [CF744A]")
	testCases := [][2]string{
		{
			`4 1 2
			1 3
			1 2`,
			`2`,
		},
		{
			`3 3 1
			2
			1 2
			1 3
			2 3`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF744A, testCases, 0)
}
