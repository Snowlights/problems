//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/451/E
// https://codeforces.com/problemset/status/451/problem/E
func TestCF451E(t *testing.T) {
	t.Log("Current test is [CF451E]")
	testCases := [][2]string{
		{
			`2 3
			1 3
			`,
			`2`,
		},
		{
			`2 4
			2 2
			`,
			`1`,
		},
		{
			`3 5
			1 3 2
			`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF451E, testCases, 0)
}
