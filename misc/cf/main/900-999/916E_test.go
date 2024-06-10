//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/916/E
// https://codeforces.com/problemset/status/916/problem/E
func TestCF916E(t *testing.T) {
	t.Log("Current test is [CF916E]")
	testCases := [][2]string{
		{
			`6 7
			1 4 2 8 5 7
			1 2
			3 1
			4 3
			4 5
			3 6
			3 1
			2 4 6 3
			3 4
			1 6
			2 2 4 -5
			1 4
			3 3
			`,
			`27
			19
			5
			`,
		},
		{
			`4 6
			4 3 5 6
			1 2
			2 3
			3 4
			3 1
			1 3
			2 2 4 3
			1 1
			2 2 4 -3
			3 1
			`,
			`18
			21
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF916E, testCases, 0)
}
