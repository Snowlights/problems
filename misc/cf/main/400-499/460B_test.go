//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/460/B
// https://codeforces.com/problemset/status/460/problem/B
func TestCF460B(t *testing.T) {
	t.Log("Current test is [CF460B]")
	testCases := [][2]string{
		{
			`3 2 8`,
			`3
			10 2008 13726`,
		},
		{
			`1 2 -18`,
			`0`,
		},
		{
			`2 2 -1`,
			`4
			1 31 337 967 `,
		},
	}
	codeforces.AssertEqualStringCase(t, CF460B, testCases, 0)
}
