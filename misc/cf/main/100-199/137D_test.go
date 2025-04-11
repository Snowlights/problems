//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/137/D
// https://codeforces.com/problemset/status/137/problem/D
func TestCF137D(t *testing.T) {
	t.Log("Current test is [CF137D]")
	testCases := [][2]string{
		{
			`abacaba
			1`,
			`0
			abacaba`,
		},
		{
			`abdcaba
			2`,
			`1
			abdcdba`,
		},
		{
			`abdcaba
			5`,
			`0
			a+b+d+c+aba`,
		},
		{
			`abacababababbcbabcd
			3`,
			`1
			abacaba+babab+bcbabcb`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF137D, testCases, 0)
}
