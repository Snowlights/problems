//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2045/A
// https://codeforces.com/problemset/status/2045/problem/A
func TestCF2045A(t *testing.T) {
	t.Log("Current test is [CF2045A]")
	testCases := [][2]string{
		{
			`ICPCJAKARTA`,
			`9`,
		},
		{
			`NGENG`,
			`5`,
		},
		{
			`YYY`,
			`3`,
		},
		{
			`DANGAN`,
			`6`,
		},
		{
			`AEIOUY`,
			`0`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF2045A, testCases, 0)
}
