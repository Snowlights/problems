//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1547/G
// https://codeforces.com/problemset/status/1547/problem/G
func TestCF1547G(t *testing.T) {
	t.Log("Current test is [CF1547G]")
	testCases := [][2]string{
		{
			`5

			6 7
			1 4
			1 3
			3 4
			4 5
			2 1
			5 5
			5 6
			
			1 0
			
			3 3
			1 2
			2 3
			3 1
			
			5 0
			
			4 4
			1 2
			2 3
			1 4
			4 3
			`,
			`1 0 1 2 -1 -1 
			1 
			-1 -1 -1 
			1 0 0 0 0 
			1 1 2 1 
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1547G, testCases, 0)
}
