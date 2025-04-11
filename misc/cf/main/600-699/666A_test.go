//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/666/A
// https://codeforces.com/problemset/status/666/problem/A
func TestCF666A(t *testing.T) {
	t.Log("Current test is [CF666A]")
	testCases := [][2]string{
		{
			`abacabaca`,
			`3
			aca
			ba
			ca`,
		},
		{
			`abaca`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF666A, testCases, 0)
}
