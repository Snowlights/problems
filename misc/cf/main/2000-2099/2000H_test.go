//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2000/H
// https://codeforces.com/problemset/status/2000/problem/H
func TestCF2000H(t *testing.T) {
	t.Log("Current test is [CF2000H]")
	testCases := [][2]string{
		{
			`3
			5
			1 2 5 905 2000000
			15
			- 2
			? 2
			? 1
			- 1
			? 1
			+ 4
			+ 2
			? 2
			+ 6
			- 4
			+ 7
			? 2
			? 3
			? 4
			? 2000000
			5
			3 4 5 6 8
			9
			? 5
			- 5
			? 5
			+ 1
			? 2
			- 6
			- 8
			+ 6
			? 5
			5
			6 7 8 9 10
			10
			? 5
			- 6
			? 4
			- 10
			+ 5
			- 8
			+ 3
			+ 2
			- 3
			+ 10`,
			`2 2 1 6 3 8 8 2000001 
			9 9 9 7 
			1 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2000H, testCases, 0)
}
