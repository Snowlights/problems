//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/23/C
// https://codeforces.com/problemset/status/23/problem/C
func TestCF23C(t *testing.T) {
	t.Log("Current test is [CF23C]")
	testCases := [][2]string{
		{
			`2
			2
			10 15
			5 7
			20 18
			1
			0 0`,
			`YES
			1 3
			YES
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF23C, testCases, 0)
}
