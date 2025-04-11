//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1844/C
// https://codeforces.com/problemset/status/1844/problem/C
func TestCF1844C(t *testing.T) {
	t.Log("Current test is [CF1844C]")
	testCases := [][2]string{
		{
			`3
			6
			-3 1 4 -1 5 -9
			5
			998244353 998244353 998244353 998244353 998244353
			1
			-2718`,
			`9
			2994733059
			-2718`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1844C, testCases, 0)
}
