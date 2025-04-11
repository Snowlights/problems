//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1269/B
// https://codeforces.com/problemset/status/1269/problem/B
func TestCF1269B(t *testing.T) {
	t.Log("Current test is [CF1269B]")
	testCases := [][2]string{
		{
			`4 3
			0 0 2 1
			2 0 1 1`,
			`1`,
		},
		{
			`3 2
			0 0 0
			1 1 1`,
			`1`,
		},
		{
			`5 10
			0 0 0 1 2
			2 1 0 0 0`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1269B, testCases, 0)
}
