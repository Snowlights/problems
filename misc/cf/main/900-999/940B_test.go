//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/940/B
// https://codeforces.com/problemset/status/940/problem/B
func TestCF940B(t *testing.T) {
	t.Log("Current test is [CF940B]")
	testCases := [][2]string{
		{
			`9
			2
			3
			1`,
			`6`,
		},
		{
			`5
			5
			2
			20
			`,
			`8`,
		},
		{
			`19
			3
			4
			2
			`,
			`12`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF940B, testCases, 0)
}
