//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1994/E
// https://codeforces.com/problemset/status/1994/problem/E
func TestCF1994E(t *testing.T) {
	t.Log("Current test is [CF1994E]")
	testCases := [][2]string{
		{
			`3
			1
			1
			
			
			2
			4
			1 2 2
			6
			1 1 3 1 3
			1
			10
			1 2 2 1 1 5 7 6 4`,
			`1
			7
			10`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1994E, testCases, 0)
}
