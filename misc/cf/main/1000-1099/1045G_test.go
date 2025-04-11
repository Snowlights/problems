//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1045/G
// https://codeforces.com/problemset/status/1045/problem/G
func TestCF1045G(t *testing.T) {
	t.Log("Current test is [CF1045G]")
	testCases := [][2]string{
		{
			`3 2
			3 6 1
			7 3 10
			10 5 8`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1045G, testCases, 0)
}
