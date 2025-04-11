//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/223/B
// https://codeforces.com/problemset/status/223/problem/B
func TestCF223B(t *testing.T) {
	t.Log("Current test is [CF223B]")
	testCases := [][2]string{
		{
			`abab
			ab`,
			`Yes`,
		},
		{
			`abacaba
			aba`,
			`No`,
		},
		{
			`abc
			ba`,
			`No`,
		},
		{
			`ababa
			abba`,
			`No`,
		},
		{
			`a
			aa`,
			`No`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF223B, testCases, 0)
}
