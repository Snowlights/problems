//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/895/C
// https://codeforces.com/problemset/status/895/problem/C
func TestCF895C(t *testing.T) {
	t.Log("Current test is [CF895C]")
	testCases := [][2]string{
		{
			`4
			1 1 1 1
			`,
			`15`,
		},
		{
			`4
			2 2 2 2
			`,
			`7`,
		},
		{
			`5
			1 2 4 5 8
			`,
			`7`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF895C, testCases, 0)
}
