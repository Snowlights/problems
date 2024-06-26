//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/620/E
// https://codeforces.com/problemset/status/620/problem/E
func TestCF620E(t *testing.T) {
	t.Log("Current test is [CF620E]")
	testCases := [][2]string{
		{
			`7 10
			1 1 1 1 1 1 1
			1 2
			1 3
			1 4
			3 5
			3 6
			3 7
			1 3 2
			2 1
			1 4 3
			2 1
			1 2 5
			2 1
			1 6 4
			2 1
			2 2
			2 3
			`,
			`2
			3
			4
			5
			1
			2
			`,
		},
		{
			`23 30
			1 2 2 6 5 3 2 1 1 1 2 4 5 3 4 4 3 3 3 3 3 4 6
			1 2
			1 3
			1 4
			2 5
			2 6
			3 7
			3 8
			4 9
			4 10
			4 11
			6 12
			6 13
			7 14
			7 15
			7 16
			8 17
			8 18
			10 19
			10 20
			10 21
			11 22
			11 23
			2 1
			2 5
			2 6
			2 7
			2 8
			2 9
			2 10
			2 11
			2 4
			1 12 1
			1 13 1
			1 14 1
			1 15 1
			1 16 1
			1 17 1
			1 18 1
			1 19 1
			1 20 1
			1 21 1
			1 22 1
			1 23 1
			2 1
			2 5
			2 6
			2 7
			2 8
			2 9
			2 10
			2 11
			2 4
			`,
			`6
			1
			3
			3
			2
			1
			2
			3
			5
			5
			1
			2
			2
			1
			1
			1
			2
			3
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF620E, testCases, 0)
}
