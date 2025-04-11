//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/150/B
// https://codeforces.com/problemset/status/150/problem/B
func TestCF150B(t *testing.T) {
	t.Log("Current test is [CF150B]")
	testCases := [][2]string{
		{
			`1 1 1`,
			`1`,
		},
		{
			`5 2 4`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF150B, testCases, 0)
}
