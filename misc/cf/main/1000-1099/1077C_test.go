//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1077/C
// https://codeforces.com/problemset/status/1077/problem/C
func TestCF1077C(t *testing.T) {
	t.Log("Current test is [CF1077C]")
	testCases := [][2]string{
		{
			`5
			2 5 1 2 2`,
			`3
			4 1 5`,
		},
		{
			`4
			8 3 5 2`,
			`2
			1 4 `,
		},
		{
			`5
			2 1 2 4 3`,
			`0

			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1077C, testCases, 0)
}
