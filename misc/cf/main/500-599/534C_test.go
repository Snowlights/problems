//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/534/C
// https://codeforces.com/problemset/status/534/problem/C
func TestCF534C(t *testing.T) {
	t.Log("Current test is [CF534C]")
	testCases := [][2]string{
		{
			`2 8
			4 4
			`,
			`3 3`,
		},
		{
			`1 3
			5
			`,
			`4`,
		},
		{
			`2 3
			2 3
			`,
			`0 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF534C, testCases, 0)
}
