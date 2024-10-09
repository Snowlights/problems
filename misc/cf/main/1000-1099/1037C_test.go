//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1037/C
// https://codeforces.com/problemset/status/1037/problem/C
func TestCF1037C(t *testing.T) {
	t.Log("Current test is [CF1037C]")
	testCases := [][2]string{
		{
			`3
			100
			001`,
			`2`,
		},
		{
			`4
			0101
			0011`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1037C, testCases, 0)
}
