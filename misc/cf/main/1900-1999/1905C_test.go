//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1905/C
// https://codeforces.com/problemset/status/1905/problem/C
func TestCF1905C(t *testing.T) {
	t.Log("Current test is [CF1905C]")
	testCases := [][2]string{
		{
			`6
			5
			aaabc
			3
			acb
			3
			bac
			4
			zbca
			15
			czddeneeeemigec
			13
			cdefmopqsvxzz`,
			`0
			1
			-1
			2
			6
			0
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1905C, testCases, 0)
}
