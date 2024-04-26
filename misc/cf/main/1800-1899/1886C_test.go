//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1886/problem/C
// https://codeforces.com/problemset/status/1886/problem/C
func TestCF1886C(t *testing.T) {
	t.Log("Current test is [CF1886C]")
	testCases := [][2]string{
		{
			`3
			cab
			6
			abcd
			9
			x
			1`,
			`abx`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1886C, testCases, 0)
}
