//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1547/E
// https://codeforces.com/problemset/status/1547/problem/E
func TestCF1547E(t *testing.T) {
	t.Log("Current test is [CF1547E]")
	testCases := [][2]string{
		{
			`5

			6 2
			2 5
			14 16
			
			10 1
			7
			30
			
			5 5
			3 1 4 2 5
			3 1 4 2 5
			
			7 1
			1
			1000000000
			
			6 3
			6 1 3
			5 5 5
			`,
			`15 14 15 16 16 17 
			36 35 34 33 32 31 30 31 32 33 
			1 2 3 4 5 
			1000000000 1000000001 1000000002 1000000003 1000000004 1000000005 1000000006 
			5 6 5 6 6 5 
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1547E, testCases, 0)
}
