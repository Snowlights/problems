//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/474/F
// https://codeforces.com/problemset/status/474/problem/F
func TestCF474F(t *testing.T) {
	t.Log("Current test is [CF474F]")
	testCases := [][2]string{
		{
			`5
			1 3 2 4 2
			4
			1 5
			2 5
			3 5
			4 5`,
			`4
			4
			1
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF474F, testCases, 0)
}
