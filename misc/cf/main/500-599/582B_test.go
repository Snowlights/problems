//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/582/B
// https://codeforces.com/problemset/status/582/problem/B
func TestCF582B(t *testing.T) {
	t.Log("Current test is [CF582B]")
	testCases := [][2]string{
		{
			`4 3
			3 1 4 2`,
			`5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF582B, testCases, 0)
}