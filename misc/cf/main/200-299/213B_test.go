//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/213/B
// https://codeforces.com/problemset/status/213/problem/B
func TestCF213B(t *testing.T) {
	t.Log("Current test is [CF213B]")
	testCases := [][2]string{
		{
			`1
			0 0 0 0 0 0 0 0 0 1`,
			`1`,
		},
		{
			`2
			1 1 0 0 0 0 0 0 0 0`,
			`1`,
		},
		{
			`3
			1 1 0 0 0 0 0 0 0 0`,
			`36`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF213B, testCases, 0)
}
