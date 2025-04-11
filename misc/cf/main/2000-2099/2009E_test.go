//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2009/E
// https://codeforces.com/problemset/status/2009/problem/E
func TestCF2009E(t *testing.T) {
	t.Log("Current test is [CF2009E]")
	testCases := [][2]string{
		{
			`4
			2 2
			7 2
			5 3
			1000000000 1000000000`,
			`1
			5
			1
			347369930`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2009E, testCases, 0)
}
