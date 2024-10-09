//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/176/B
// https://codeforces.com/problemset/status/176/problem/B
func TestCF176B(t *testing.T) {
	t.Log("Current test is [CF176B]")
	testCases := [][2]string{
		{
			`ab
			ab
			2
			`,
			`1`,
		},
		{
			`ababab
			ababab
			1
			`,
			`2`,
		},
		{
			`ab
			ba
			2
			`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF176B, testCases, 0)
}
