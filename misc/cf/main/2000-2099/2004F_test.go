//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2004/F
// https://codeforces.com/problemset/status/2004/problem/F
func TestCF2004F(t *testing.T) {
	t.Log("Current test is [CF2004F]")
	testCases := [][2]string{
		{
			`4
			3
			2 1 3
			4
			1 1 1 1
			5
			4 2 3 1 5
			4
			1 2 1 2`,
			`3
			0
			14
			5`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF2004F, testCases, 0)
}
