//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1238/E
// https://codeforces.com/problemset/status/1238/problem/E
func TestCF1238E(t *testing.T) {
	t.Log("Current test is [CF1238E]")
	testCases := [][2]string{
		{
			`6 3
			aacabc`,
			`5`,
		},
		{
			`6 4
			aaaaaa`,
			`0`,
		},
		{
			`15 4
			abacabadabacaba`,
			`16`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1238E, testCases, 0)
}
