//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1799/C
// https://codeforces.com/problemset/status/1799/problem/C
func TestCF1799C(t *testing.T) {
	t.Log("Current test is [CF1799C]")
	testCases := [][2]string{
		{
			`12
			a
			aab
			abb
			abc
			aabb
			aabbb
			aaabb
			abbb
			abbbb
			abbcc
			eaga
			ffcaba`,
			`a
			aba
			bab
			bca
			abba
			abbba
			ababa
			bbab
			bbabb
			bbcca
			agea
			acffba
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1799C, testCases, 0)
}
