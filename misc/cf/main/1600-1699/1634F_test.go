//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1634/problem/F
// https://codeforces.com/problemset/status/1634/problem/F
func TestCF1634F(t *testing.T) {
	t.Log("Current test is [CF1634F]")
	testCases := [][2]string{
		{
			`3 5 3
			2 2 1
			0 0 0
			A 1 3
			A 1 3
			B 1 1
			B 2 2
			A 3 3
			`,
			`YES
			NO
			NO
			NO
			YES`,
		},
		{
			`5 3 10
			2 5 0 3 5
			3 5 8 2 5
			B 2 3
			B 3 4
			A 1 2
			`,
			`NO
			NO
			YES
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1634F, testCases, 0)
}
