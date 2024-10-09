//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1097/C
// https://codeforces.com/problemset/status/1097/problem/C
func TestCF1097C(t *testing.T) {
	t.Log("Current test is [CF1097C]")
	testCases := [][2]string{
		{
			`7
			)())
			)
			((
			((
			(
			)
			)
			`,
			`2`,
		},
		{
			`4
			(
			((
			(((
			(())
			`,
			`0`,
		},
		{
			`2
			(())
			()
			`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1097C, testCases, 0)
}
