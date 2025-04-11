//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/113/B
// https://codeforces.com/problemset/status/113/problem/B
func TestCF113B(t *testing.T) {
	t.Log("Current test is [CF113B]")
	testCases := [][2]string{
		{
			`round
			ro
			ou`,
			`1`,
		},
		{
			`codeforces
			code
			forca`,
			`0`,
		},
		{
			`abababab
			a
			b`,
			`4`,
		},
		{
			`aba
			ab
			ba`,
			`1`,
		},
		{
			`aaaaaaaaa
			aa
			aaa`,
			`7`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF113B, testCases, 0)
}
