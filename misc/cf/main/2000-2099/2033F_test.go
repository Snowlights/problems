//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2033/F
// https://codeforces.com/problemset/status/2033/problem/F
func TestCF2033F(t *testing.T) {
	t.Log("Current test is [CF2033F]")
	testCases := [][2]string{
		{
			`3
			3 2
			100 1
			1000000000000 1377`,
			`9
			100
			999244007`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2033F, testCases, 0)
}
