//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/888/problem/C
// https://codeforces.com/problemset/status/888/problem/C
func TestCF888C(t *testing.T) {
	t.Log("Current test is [CF888C]")
	testCases := [][2]string{
		{
			`abacaba`,
			`2`,
		},
		{
			`zzzzz`,
			`1`,
		},
		{
			`abcde`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF888C, testCases, 0)
}
