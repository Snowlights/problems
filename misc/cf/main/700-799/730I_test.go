//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/730/I
// https://codeforces.com/problemset/status/730/problem/I
func TestCF730I(t *testing.T) {
	t.Log("Current test is [CF730I]")
	testCases := [][2]string{
		{
			`5 2 2
			1 3 4 5 2
			5 3 2 1 4
			`,
			`18
			3 4 
			1 5`,
		},
		{
			`4 2 2
			10 8 8 3
			10 7 9 4
			`,
			`31
			1 2 
			3 4 
			`,
		},
		{
			`5 3 1
			5 2 5 1 7
			6 3 1 6 3
			`,
			`23
			1 3 5 
			4 
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF730I, testCases, 0)
}
