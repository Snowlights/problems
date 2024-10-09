//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1301/F
// https://codeforces.com/problemset/status/1301/problem/F
func TestCF1301F(t *testing.T) {
	t.Log("Current test is [CF1301F]")
	testCases := [][2]string{
		{
			`3 4 5
			1 2 1 3
			4 4 5 5
			1 2 1 3
			2
			1 1 3 4
			2 2 2 2`,
			`2
			0`,
		},
		{
			`4 4 8
			1 2 2 8
			1 3 4 7
			5 1 7 6
			2 3 8 8
			4
			1 1 2 2
			1 1 3 4
			1 1 2 4
			1 1 4 4`,
			`2
			3
			3
			4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1301F, testCases, 0)
}
