//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1492/C
// https://codeforces.com/problemset/status/1492/problem/C
func TestCF1492C(t *testing.T) {
	t.Log("Current test is [CF1492C]")
	testCases := [][2]string{
		{
			`5 3
			abbbc
			abc`,
			`3`,
		},
		{
			`5 2
			aaaaa
			aa`,
			`4`,
		},
		{
			`5 5
			abcdf
			abcdf`,
			`1`,
		},
		{
			`2 2
			ab
			ab`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1492C, testCases, 0)
}
