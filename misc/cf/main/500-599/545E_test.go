//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/545/E
// https://codeforces.com/problemset/status/545/problem/E
func TestCF545E(t *testing.T) {
	t.Log("Current test is [CF545E]")
	testCases := [][2]string{
		{
			`3 3
			1 2 1
			2 3 1
			1 3 2
			3`,
			`2
			1 2`,
		},
		{
			`4 4
			1 2 1
			2 3 1
			3 4 1
			4 1 2
			4`,
			`4
			2 3 4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF545E, testCases, 0)
}
