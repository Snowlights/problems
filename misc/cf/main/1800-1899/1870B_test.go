//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1870/B
// https://codeforces.com/problemset/status/1870/problem/B
func TestCF1870B(t *testing.T) {
	t.Log("Current test is [CF1870B]")
	testCases := [][2]string{
		{
			`2
			2 3
			0 1
			1 2 3
			3 1
			1 1 2
			1`,
			`0 1
			2 3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1870B, testCases, 0)
}
