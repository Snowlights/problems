//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1607/D
// https://codeforces.com/problemset/status/1607/problem/D
func TestCF1607D(t *testing.T) {
	t.Log("Current test is [CF1607D]")
	testCases := [][2]string{
		{
			`8
			4
			1 2 5 2
			BRBR
			2
			1 1
			BB
			5
			3 1 4 2 5
			RBRRB
			5
			3 1 3 1 3
			RBRRB
			5
			5 1 5 1 5
			RBRRB
			4
			2 2 2 2
			BRBR
			2
			1 -2
			BR
			4
			-2 -1 4 0
			RRRR`,
			`YES
			NO
			YES
			YES
			NO
			YES
			YES
			YES
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1607D, testCases, 0)
}
