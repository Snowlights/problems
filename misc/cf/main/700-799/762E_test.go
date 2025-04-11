//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/762/E
// https://codeforces.com/problemset/status/762/problem/E
func TestCF762E(t *testing.T) {
	t.Log("Current test is [CF762E]")
	testCases := [][2]string{
		{
			`3 2
			1 3 10
			3 2 5
			4 10 8`,
			`1`,
		},
		{
			`3 3
			1 3 10
			3 2 5
			4 10 8`,
			`2`,
		},
		{
			`5 1
			1 3 2
			2 2 4
			3 2 1
			4 2 1
			5 3 3`,
			`2`,
		},
		{
			`5 1
			1 5 2
			2 5 4
			3 5 1
			4 5 1
			5 5 3`,
			`5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF762E, testCases, 0)
}
