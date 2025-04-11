//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2074/C
// https://codeforces.com/problemset/status/2074/problem/C
func TestCF2074C(t *testing.T) {
	t.Log("Current test is [CF2074C]")
	testCases := [][2]string{
		{
			`7
			5
			2
			6
			3
			69
			4
			420`,
			`3
			-1
			5
			-1
			66
			-1
			320`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2074C, testCases, 0)
}
