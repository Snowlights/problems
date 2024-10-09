//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1468/J
// https://codeforces.com/problemset/status/1468/problem/J
func TestCF1468J(t *testing.T) {
	t.Log("Current test is [CF1468J]")
	testCases := [][2]string{
		{
			`4
			4 5 7
			4 1 3
			1 2 5
			2 3 8
			2 4 1
			3 4 4
			4 6 5
			1 2 1
			1 3 1
			1 4 2
			2 4 1
			4 3 1
			3 2 1
			3 2 10
			1 2 8
			1 3 10
			5 5 15
			1 2 17
			3 1 15
			2 3 10
			1 4 14
			2 5 8`,
			`1
			3
			0
			0`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1468J, testCases, 0)
}
