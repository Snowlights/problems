//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2037/F
// https://codeforces.com/problemset/status/2037/problem/F
func TestCF2037F(t *testing.T) {
	t.Log("Current test is [CF2037F]")
	testCases := [][2]string{
		{
			`6
			5 5 3
			7 7 7 7 7
			1 2 3 4 5
			9 5 9
			2 4 6 8 10 8 6 4 2
			1 2 3 4 5 6 7 8 9
			2 10 2
			1 1
			1 20
			2 10 1
			69696969 420420420
			1 20
			2 10 2
			10 15
			1 19
			2 2 2
			1000000000 1
			1 3`,
			`2
			2
			-1
			6969697
			15
			1000000000`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2037F, testCases, 0)
}
