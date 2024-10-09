//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1362/C
// https://codeforces.com/problemset/status/1362/problem/C
func TestCF1362C(t *testing.T) {
	t.Log("Current test is [CF1362C]")
	testCases := [][2]string{
		{
			`5
			5
			7
			11
			1
			2000000000000
			`,
			`8
			11
			19
			1
			3999999999987
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1362C, testCases, 0)
}
