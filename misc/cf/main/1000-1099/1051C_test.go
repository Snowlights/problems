//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1051/C
// https://codeforces.com/problemset/status/1051/problem/C
func TestCF1051C(t *testing.T) {
	t.Log("Current test is [CF1051C]")
	testCases := [][2]string{
		{
			`4
			3 5 7 1
			`,
			`YES
			BABA
			`,
		},
		{
			`3
			3 5 1
			`,
			`NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1051C, testCases, 0)
}
