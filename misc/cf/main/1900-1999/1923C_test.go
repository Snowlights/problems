//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1923/C
// https://codeforces.com/problemset/status/1923/problem/C
func TestCF1923C(t *testing.T) {
	t.Log("Current test is [CF1923C]")
	testCases := [][2]string{
		{
			`1
			5 4
			1 2 1 4 5
			1 5
			4 4
			3 4
			1 3`,
			`YES
			NO
			YES
			NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1923C, testCases, 0)
}
