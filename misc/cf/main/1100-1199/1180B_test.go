//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1180/B
// https://codeforces.com/problemset/status/1180/problem/B
func TestCF1180B(t *testing.T) {
	t.Log("Current test is [CF1180B]")
	testCases := [][2]string{
		{
			`4
			2 2 2 2
			`,
			`-3 -3 -3 -3`,
		},
		{
			`1
			0
			`,
			`0`,
		},
		{
			`3
			-3 -3 2
			`,
			`-3 -3 2 `,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1180B, testCases, 0)
}
