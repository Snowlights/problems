//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2014/F
// https://codeforces.com/problemset/status/2014/problem/F
func TestCF2014F(t *testing.T) {
	t.Log("Current test is [CF2014F]")
	testCases := [][2]string{
		{
			`5
			3 1
			2 3 1
			1 2
			2 3
			3 1
			3 6 3
			1 2
			2 3
			3 1
			-2 -3 -1
			1 2
			2 3
			6 1
			5 -4 3 6 7 3
			4 1
			5 1
			3 5
			3 6
			1 2
			8 1
			3 5 2 7 8 5 -3 -4
			7 3
			1 8
			4 3
			3 5
			7 6
			8 7
			2 1`,
			`3
			8
			0
			17
			26`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2014F, testCases, 0)
}
