//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1906/H
// https://codeforces.com/problemset/status/1906/problem/H
func TestCF1906H(t *testing.T) {
	t.Log("Current test is [CF1906H]")
	testCases := [][2]string{
		{
			`3 4
			AMA
			ANAB`,
			`9`,
		},
		{
			`5 8
			BINUS
			BINANUSA`,
			`120`,
		},
		{
			`15 30
			BINUSUNIVERSITY
			BINANUSANTARAUNIVERSITYJAKARTA`,
			`151362308`,
		},
		{
			`4 4
			UDIN
			ASEP`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1906H, testCases, 0)
}
