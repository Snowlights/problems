//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/435/B
// https://codeforces.com/problemset/status/435/problem/B
func TestCF435B(t *testing.T) {
	t.Log("Current test is [CF435B]")
	testCases := [][2]string{
		{
			`1990 1`,
			`9190`,
		},
		{
			`300 0`,
			`300`,
		},
		{
			`1034 2`,
			`3104`,
		},
		{
			`9090000078001234 6`,
			`9907000008001234`,
		},
		{
			`9022 2`,
			`9220`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF435B, testCases, 0)
}
