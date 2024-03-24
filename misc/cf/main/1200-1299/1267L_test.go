//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1267/L
// https://codeforces.com/problemset/status/1267/problem/L
func TestCF1267L(t *testing.T) {
	t.Log("Current test is [CF1267L]")
	testCases := [][2]string{
		{
			`3 2 2
			abcdef
			`,
			`af
			bc
			ed`,
		},
		{
			`2 3 1
			abcabc
			`,
			`aab
			bcc
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1267L, testCases, 0)
}
