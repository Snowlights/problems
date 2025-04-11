//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1795/C
// https://codeforces.com/problemset/status/1795/problem/C
func TestCF1795C(t *testing.T) {
	t.Log("Current test is [CF1795C]")
	testCases := [][2]string{
		{
			`4
			3
			10 20 15
			9 8 6
			1
			5
			7
			4
			13 8 5 4
			3 4 2 1
			3
			1000000000 1000000000 1000000000
			1 1 1000000000`,
			`9 9 12 
			5 
			3 8 6 4 
			1 2 2999999997`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1795C, testCases, 0)
}
