//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/156/C
// https://codeforces.com/problemset/status/156/problem/C
func TestCF156C(t *testing.T) {
	t.Log("Current test is [CF156C]")
	testCases := [][2]string{
		{
			`1
			ab`,
			`1`,
		},
		{
			`1
			aaaaaaaaaaa`,
			`0`,
		},
		{
			`2
			ya
			klmbfxzb`,
			`24
			320092793`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF156C, testCases, 0)
}
