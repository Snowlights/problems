//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/990/C
// https://codeforces.com/problemset/status/990/problem/C
func TestCF990C(t *testing.T) {
	t.Log("Current test is [CF990C]")
	testCases := [][2]string{
		{
			`3
			)
			()
			(`,
			`2`,
		},
		{
			`2
			()
			()`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF990C, testCases, 0)
}
