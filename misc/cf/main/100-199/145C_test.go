//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/145/C
// https://codeforces.com/problemset/status/145/problem/C
func TestCF145C(t *testing.T) {
	t.Log("Current test is [CF145C]")
	testCases := [][2]string{
		{
			`3 2
			10 10 10`,
			`3`,
		},
		{
			`4 2
			4 4 7 7`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF145C, testCases, 0)
}
