//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/460/C
// https://codeforces.com/problemset/status/460/problem/C
func TestCF460C(t *testing.T) {
	t.Log("Current test is [CF460C]")
	testCases := [][2]string{
		{
			`6 2 3
			2 2 2 2 1 1`,
			`2`,
		},
		{
			`2 5 1
			5 8`,
			`9`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF460C, testCases, 0)
}
