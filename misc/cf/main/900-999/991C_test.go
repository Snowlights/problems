//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/991/C
// https://codeforces.com/problemset/status/991/problem/C
func TestCF991C(t *testing.T) {
	t.Log("Current test is [CF991C]")
	testCases := [][2]string{
		{
			`68`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF991C, testCases, 0)
}
