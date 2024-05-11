//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1299/A
// https://codeforces.com/problemset/status/1299/problem/A
func TestCF1299A(t *testing.T) {
	t.Log("Current test is [CF1299A]")
	testCases := [][2]string{
		{
			`4
			4 0 11 6
			`,
			`11 6 4 0 `,
		},
		{
			`1
			13
			`,
			`13`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1299A, testCases, 0)
}
