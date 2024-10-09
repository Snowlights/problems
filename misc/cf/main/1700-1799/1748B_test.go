//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1748/B
// https://codeforces.com/problemset/status/1748/problem/B
func TestCF1748B(t *testing.T) {
	t.Log("Current test is [CF1748B]")
	testCases := [][2]string{
		{
			`7
			1
			7
			2
			77
			4
			1010
			5
			01100
			6
			399996
			5
			23456
			18
			789987887987998798`,
			`1
			2
			10
			12
			10
			15
			106
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1748B, testCases, 0)
}
