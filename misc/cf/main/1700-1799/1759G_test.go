//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1759/G
// https://codeforces.com/problemset/status/1759/problem/G
func TestCF1759G(t *testing.T) {
	t.Log("Current test is [CF1759G]")
	testCases := [][2]string{
		{
			`6
			6
			4 3 6
			4
			2 4
			8
			8 7 2 3
			6
			6 4 2
			4
			4 4
			8
			8 7 4 5`,
			`1 4 2 3 5 6 
			1 2 3 4 
			-1
			5 6 3 4 1 2 
			-1
			1 8 6 7 2 4 3 5 
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1759G, testCases, 0)
}
