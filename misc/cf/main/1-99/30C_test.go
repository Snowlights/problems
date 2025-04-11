//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/30/C
// https://codeforces.com/problemset/status/30/problem/C
func TestCF30C(t *testing.T) {
	t.Log("Current test is [CF30C]")
	testCases := [][2]string{
		{
			`1
			0 0 0 0.5`,
			`0.5000000000`,
		},
		{
			`2
			0 0 0 0.6
			5 0 5 0.7`,
			`1.3000000000`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF30C, testCases, 0)
}
