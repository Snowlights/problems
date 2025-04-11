//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2023/A
// https://codeforces.com/problemset/status/2023/problem/A
func TestCF2023A(t *testing.T) {
	t.Log("Current test is [CF2023A]")
	testCases := [][2]string{
		{
			`4
			2
			1 4
			2 3
			3
			3 2
			4 3
			2 1
			5
			5 10
			2 3
			9 6
			4 1
			8 7
			1
			10 20`,
			`2 3 1 4
			2 1 3 2 4 3
			4 1 2 3 5 10 8 7 9 6
			10 20`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2023A, testCases, 0)
}
