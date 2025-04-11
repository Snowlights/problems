//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1948/D
// https://codeforces.com/problemset/status/1948/problem/D
func TestCF1948D(t *testing.T) {
	t.Log("Current test is [CF1948D]")
	testCases := [][2]string{
		{
			`4
			zaabaabz
			?????
			code?????s
			codeforces`,
			`6
			4
			10
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1948D, testCases, 0)
}
