//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/17/E
// https://codeforces.com/problemset/status/17/problem/E
func TestCF17E(t *testing.T) {
	t.Log("Current test is [CF17E]")
	testCases := [][2]string{
		{
			`4
			babb`,
			`6`,
		},
		{
			`2
			aa`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF17E, testCases, 0)
}
