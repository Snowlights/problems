//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1822/E
// https://codeforces.com/problemset/status/1822/problem/E
func TestCF1822E(t *testing.T) {
	t.Log("Current test is [CF1822E]")
	testCases := [][2]string{
		{
			`10
			10
			codeforces
			3
			abc
			10
			taarrrataa
			10
			dcbdbdcccc
			4
			wwww
			12
			cabbaccabaac
			10
			aadaaaaddc
			14
			aacdaaaacadcdc
			6
			abccba
			12
			dcbcaebacccd
			`,
			`0
			-1
			1
			1
			-1
			3
			-1
			2
			2
			2
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1822E, testCases, 0)
}
