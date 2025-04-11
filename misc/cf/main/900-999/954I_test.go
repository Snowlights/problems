//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/954/I
// https://codeforces.com/problemset/status/954/problem/I
func TestCF954I(t *testing.T) {
	t.Log("Current test is [CF954I]")
	testCases := [][2]string{
		{
			`abcdefa
			ddcb`,
			`2 3 3 3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF954I, testCases, 0)
}
