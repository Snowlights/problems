//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/379/F
// https://codeforces.com/problemset/status/379/problem/F
func TestCF379F(t *testing.T) {
	t.Log("Current test is [CF379F]")
	testCases := [][2]string{
		{
			`5
			2
			3
			4
			8
			5`,
			`3
			4
			4
			5
			6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF379F, testCases, 0)
}
