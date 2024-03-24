//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/920/C
// https://codeforces.com/problemset/status/920/problem/C
func TestCF920C(t *testing.T) {
	t.Log("Current test is [CF920C]")
	testCases := [][2]string{
		{
			`6
			1 2 5 3 4 6
			01110
			`,
			`YES`,
		},
		{
			`6
			1 2 5 3 4 6
			01010
			`,
			`NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF920C, testCases, 0)
}
