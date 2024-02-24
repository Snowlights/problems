//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/414/problem/B
// https://codeforces.com/problemset/status/414/problem/B
func TestCF414B(t *testing.T) {
	t.Log("Current test is [CF414B]")
	testCases := [][2]string{
		{
			`3 2`,
			`5`,
		},
		{
			`6 4`,
			`39`,
		},
		{
			`2 1`,
			`2`,
		},
		{
			`1415 562`,
			`953558593`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF414B, testCases, 0)
}
