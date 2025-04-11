//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/843/B
// https://codeforces.com/problemset/status/843/problem/B
func TestCF843B(t *testing.T) {
	t.Log("Current test is [CF843B]")
	testCases := [][2]string{
		{
			`5 3 80
			97 -1
			58 5
			16 2
			81 1
			79 4`,
			`? 1
			? 2
			? 3
			? 4
			? 5
			! 81`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF843B, testCases, 0)
}
