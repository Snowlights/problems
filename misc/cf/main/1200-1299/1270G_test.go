//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1270/G
// https://codeforces.com/problemset/status/1270/problem/G
func TestCF1270G(t *testing.T) {
	t.Log("Current test is [CF1270G]")
	testCases := [][2]string{
		{
			`2
			5
			0 1 2 3 4
			4
			-3 1 1 1
			`,
			`1
			1 
			4
			1 4 3 2 
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1270G, testCases, 0)
}
