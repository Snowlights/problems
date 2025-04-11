//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2046/B
// https://codeforces.com/problemset/status/2046/problem/B
func TestCF2046B(t *testing.T) {
	t.Log("Current test is [CF2046B]")
	testCases := [][2]string{
		{
			`3
			3
			2 1 3
			5
			1 2 2 1 4
			6
			1 2 3 6 5 4`,
			`1 3 3 
			1 1 3 3 5 
			1 2 3 4 6 7`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2046B, testCases, 0)
}
