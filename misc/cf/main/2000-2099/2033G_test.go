//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2033/G
// https://codeforces.com/problemset/status/2033/problem/G
func TestCF2033G(t *testing.T) {
	t.Log("Current test is [CF2033G]")
	testCases := [][2]string{
		{
			`3
			5
			1 2
			2 3
			3 4
			3 5
			3
			5 1
			3 1
			2 0
			9
			8 1
			1 7
			1 4
			7 3
			4 9
			3 2
			1 5
			3 6
			7
			6 0
			2 3
			6 2
			8 2
			2 4
			9 2
			6 3
			6
			2 1
			2 5
			2 4
			5 6
			4 3
			3
			3 1
			1 3
			6 5`,
			`2 1 2 
			0 5 2 4 5 5 5 
			1 3 4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2033G, testCases, 0)
}
