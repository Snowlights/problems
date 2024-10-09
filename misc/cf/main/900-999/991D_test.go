//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/991/D
// https://codeforces.com/problemset/status/991/problem/D
func TestCF991D(t *testing.T) {
	t.Log("Current test is [CF991D]")
	testCases := [][2]string{
		{
			`00
			00`,
			`1`,
		},
		{
			`00X00X0XXX0
			0XXX0X00X00`,
			`4`,
		},
		{
			`0X0X0
			0X0X0`,
			`0`,
		},
		{
			`0XXX0
			00000`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF991DV2, testCases, 0)
}
